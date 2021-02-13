"""
Here's the majority of difference from unittest (I assume), mine's prettier :)

"""

import traceback

class prettyTestPrint:
	def __init__(self):
		# Colors
		self.RED   = "\x1b[1;31m"
		self.GREEN = "\x1b[92m"
		self.BLUE  = "\x1b[34m"
		self.UNSET = "\x1b[0m"
		# aggregate values
		self.testPass = 0
		self.testFail = 0
		self.assertPass = 0
		self.assertFail = 0

	def name(self, name):
		print(name)

	def passed(self, name):
		print(name + "[" + self.GREEN + "passed" + self.UNSET + "]")
		self.testPass = self.testPass + 1

	def failed(self, name):
		print(name + " [" + self.RED + "failed" + self.UNSET + "]")
		self.testFail = self.testFail + 1

	def asserts(self, a, b, op, passed):
		if passed:
			print("\t[" + self.GREEN + "ASSERT" + self.UNSET + "] ", a, op, b)
			self.assertPass = self.assertPass + 1
		else:
			print("\t[" + self.RED + "ASSERT" + self.UNSET + "] ", a, op, b)
			self.assertFail = self.assertFail + 1

	def trace(self, type, value, trace):
		print(type)
		print(value)
		traceback.print_tb(trace)

	def aggregate(self):
		print("\nYou passed", end='')
		if self.testFail == 0:
			print(self.GREEN, self.testPass, self.UNSET, end='')
		else:
			print(self.BLUE, self.testPass, self.UNSET, end='')
		print("out of" + self.GREEN, (self.testPass + self.testFail), self.UNSET + "tests\n")
