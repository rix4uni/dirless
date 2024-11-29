## dirless

dirless is a CLI tool to match, highlight, and categorize URLs using configurable regex patterns.

## Installation
```
go install github.com/rix4uni/dirless@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/dirless/releases/download/v0.0.1/dirless-linux-amd64-0.0.1.tgz
tar -xvzf dirless-linux-amd64-0.0.1.tgz
rm -rf dirless-linux-amd64-0.0.1.tgz
mv dirless ~/go/bin/dirless
```
Or download [binary release](https://github.com/rix4uni/dirless/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/dirless.git
cd dirless; go install
```

## Usage
```
Usage of dirless:
  -match string
        Path to custom match.json file.
  -nc
        Disable colors in CLI output.
  -silent
        Silent mode.
  -verbose
        Enable verbose output.
  -version
        Print the version of the tool and exit.
```

## Examples
```
▶ cat targets.txt
https://www.facebook.com/sharer.php?u=https://blog.rix4uni.com/2024/11/today-gold-and-the-others-price/management.html
https://www.rix4uni.com/admin/idea/265438/rumored-buzz-on-the-best-rv-batteries-near-me
https://www.rix4uni.com/admin/idea/359642/default?r=36458
https://dev-downloads.rix4uni.com/2024-05-18/Repository/test
https://www.rix4uni.com/browse/Flash/694582/how-to-bathe?cat=1
https://dev-downloads.rix4uni.com/how-to-bathe?cat=1
https://dev-downloads.rix4uni.com/786543/2024-05-18/Repository/test

▶ cat targets.txt | dirless -silent -verbose
[regex1] [id3951] https://www.rix4uni.com/admin/idea/265438/rumored-buzz-on-the-best-rv-batteries-near-me
[regex1] [id3951] [ignored] https://www.rix4uni.com/admin/idea/359642/default?r=36458
[regex1] [id3951] [ignored] https://www.rix4uni.com/browse/Flash/694582/how-to-bathe?cat=1
[regex2] [id7869] https://dev-downloads.rix4uni.com/786543/2024-05-18/Repository/test
[regex3] [id7869] [ignored] https://dev-downloads.rix4uni.com/2024-05-18/Repository/test
[unmatched] https://www.facebook.com/sharer.php?u=https://blog.rix4uni.com/2024/11/today-gold-and-the-others-price/management.html
[unmatched] https://dev-downloads.rix4uni.com/how-to-bathe?cat=1
```