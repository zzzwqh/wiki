def git_url = "https://gitlab.itsky.tech/demo/prometheus.git"
def trim_git_url = git_url.replaceFirst("https://", "")
 println trim_git_url
 def ENV= "test"
if (ENV.equals("test")){
def gettags = ("git ls-remote -h https://root:wqh127.0.0.1@${trim_git_url} release*").execute()
println gettags.text.readLines().collect { it.split()[1].replaceAll('refs/heads/', '') }.unique()
}
else{
def gettags = ("git ls-remote -h https://root:wqh127.0.0.1@${trim_git_url} develop feature*").execute()
println gettags.text.readLines().collect { it.split()[1].replaceAll('refs/heads/', '') }.unique()
} 