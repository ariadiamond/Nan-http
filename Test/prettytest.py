"""
This is very inspired by the unittest module, but I was not happy with the
aesthetics, so I decided to write my own.

"""

import sys
from prettyTestPrint import prettyTestPrint

# Verbosity Levels
# 0 - assertions exceptions fail and return control
# 1 - failed assertion alert, but flow continues
# 2 - all assertions are printed, but assertions return control
# 3 - all assertions are printed, but flow continues

class prettyTest:
	def __init__(self, tests={}, verbosity=2):
		self.tests   = tests
		self.verbose = verbosity # Get verbosity here?
		self.pretty  = prettyTestPrint()

	def main(self): # dictionaries
		for i in list(self.tests):
			self.runTest(i, self.tests[i][0], self.tests[i][1:])

		self.pretty.aggregate()

	def runTest(self, name, f, args):
		try:
			if len(args) == 0:
				f()
			elif len(args) == 1:
				f(args[0])
			elif len(args) == 2:
				f(args[0], args[1])
			elif len(args) == 3:
				f(args[0], args[1], args[2])
			elif len(args) == 4:
				f(args[0], args[1], args[2], args[3])
		except AttributeError:
			print("hello")
			return
		except AssertionError:
			self.pretty.failed(name)
			type, val, trace = sys.exc_info()
			self.pretty.trace(type, val, trace)
			if (self.verbose & 0x01) == 1: # just move onto the next test
				return
			else:
				self.pretty.aggregate()
				sys.exit(1)
		except:
			type, val, trace = sys.exc_info()
			self.pretty.trace(type, val, trace)
			self.pretty.aggregate()
			sys.exit(1)
		self.pretty.passed(name)

	def asserts(self, a, b, op):
		assertion = False
		if op == '==':
			assertion = a == b
		elif op == '!=':
			assertion = a != b
		elif op == '<':
			assertion = a < b
		elif op == '<=':
			assertion = a <= b
		elif op == '>':
			assertion = a > b
		elif op == '>=':
			assertion = a >= b
		# elif op == 'in': # do in and not in D.R.Y?
		# else is to implicitly fail (not implmented)

		# print results
		if assertion:
			if self.verbose >= 2:
				self.pretty.asserts(a, b, op, True)
		else:
			self.pretty.asserts(a, b, op, False)
			raise AssertionError

	def assertFile(self, content, filename, isBinary=False):
		try:
			if isBinary:
				f = open(filename, "rb")
			else:
				f = open(filename)
		except OSError:
			self.pretty.asserts("Failed to open:", "", filename, False)
			return
		file = f.read()
		if file in content: # found it
			if self.verbose >= 2:
				self.pretty.asserts("File:", filename, "", True)
		else: # did not find it
			self.pretty.asserts("File:", filename, "", False)
			raise AssertionError
