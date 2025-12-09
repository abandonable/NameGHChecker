<p align="center">
<h1>Github Name Checker</h1>
</p>

## Information about the Project.
> This was just kinda a on the fly project to get a sick name on my new Github Account <br>I used a English Dictionary Wordlist Provided by [dwyl](https://github.com/dwyl/english-words) and his cool EN Dictionary Lists

## How to use my GH Name Checker?

```
git clone https://github.com/abandonable/NameGHChecker
cd ./NameGHChecker

wget https://raw.githubusercontent.com/dwyl/english-words/refs/heads/master/words_alpha.txt -o ./words.txt 
# If you like to get the list i used ^^^^^^^

go run . 
# Or "go build ." if you want a Static binary directly
```

## How can i Customize my experience?
> Just Change the Default config `config.toml` which is in the Project root.

```toml
Wordlist = "./words.txt"
User_Agent = "Mozilla/5.0 (X11; Linux x86_64; rv:145.0) Gecko/20100101 Firefox/145.0"
Ratelimit = 120
RequestDelay = 1
```
