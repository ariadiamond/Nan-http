import requests
from prettytest import prettyTest

class tests(prettyTest):
	def __init__(self, url = "http://localhost:8085/"):
		super().__init__(verbosity=0)
		self.url = url
		self.tests = {
			"Get Default Stylesheet" : [self.file, "style.css", "../../Root/style.css"],
			"Get favicon.ico"        : [self.file, "favicon.ico", "../../Root/favicon.ico", True],
			"Get favicon.png"        : [self.file, "favicon.png", "../../Root/favicon.png", True],
			"Get apple favicon"      : [self.file, "apple-touch-icon.png", "../../Root/favicon.png", True],
			"Get Home Index"         : [self.construct, "index", ["../../Root/index.html"]],
		}

	# path is the url path for the server and compare it to file
	def file(self, path, file, isBinary=False):
		resp = requests.get(self.url + path)
		if resp.status_code != 200:
			raise AttributeError
		if isBinary:
			self.assertFile(resp.content, file, True)
		else:
			self.assertFile(resp.text, file)

	# Constructed pages do not support binary yet
	def construct(self, path, files):
		files += ["../../Root/footer.html"]
		resp = requests.get(self.url + path)
		if resp.status_code != 200:
			raise AttributeError
		for i in files:
			self.assertFile(resp.text, i)


	def fileNotFound(self, path):
		resp = requests.get(self.url + path)
		self.asserts(resp.status_code, 404, "==")

	def forbidden(self, path):
		resp = requests.get(self.url + path)
		self.asserts(resp.status_code, 403, "==")

if __name__ == "__main__":
	test = tests()
	test.main()
