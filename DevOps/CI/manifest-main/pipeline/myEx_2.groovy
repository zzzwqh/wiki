def repoUrl = SelectRepo.replaceFirst("https://", "")
def cmd = """
git ls-remote --heads https://oauth2:J9A2wN6qmHZwRSd1miMN@$repoUrl
"""
def process = cmd.execute()
def output = process.text.trim()
def regex = /refs\/heads\/([\w.-]+)/
def branches = []
(output =~ regex).each { match ->
    branches.add(match[1])
}
 
return branches 