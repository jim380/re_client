# Re Client

Re_client is a Command Line Interface (CLI) application designed to convert FIX messages from [RE](https://github.com/jim380/Re) to the Financial Information eXchange (FIX) format

# Installation
Pre-requisites

[Go 1.16+](https://go.dev/doc/install)

A Unix/Linux terminal (Mac, Ubuntu, CentOS, etc.)

# Build from Source
Clone the repository:
```
git clone https://github.com/jim380/re_client.git
```
Navigate to the cloned directory and build the application:
```
cd re_client && make build
```

Install the Application

Run the following command to install the executable:
```
sudo make install
```
Usage

Run Re_client in your terminal to see the list of available commands and options:
```
Re_client
```

Available Commands
```
 Re_client order [address] [flags]
```
```
 Re_client orders [flags]
```
```
 Re_client execution-report [address] [flags]
```
```
 Re_client execution-reports [flags]
```

# [Sample of generated FIX messages](https://github.com/jim380/re_client/tree/main/RawFixMessages)
 8=FIX4.4|9=0|35=D|49=cosmos1nvcgd368m4pm5mm3ppzawhsq6grra4ejnppplx|56=cosmosvaloper1zqgheeawp7cmqk27dgyctd80rd8ryhqs6la9wc|34=16801684|52=2023-08-31T10:19:49Z|11=07C038FC655EDDEBE0F5E99BAEB05E47FF2DD5FCFE422C4B14D0DF5078417693|55=uatom|54=2|44=117401|59=1|

 8=FIX4.4|9=0|35=D|49=cosmos1tfrmhsw9hymfe7vfcxaly8xvrqtaaaxm9648ry|56=cosmosvaloper1zqgheeawp7cmqk27dgyctd80rd8ryhqs6la9wc|34=16801845|52=2023-08-31T10:36:14Z|11=08EC92BF1B03C1C5E07412B0254203C76118FA97F2B6CDBD91CCD62F02C5C5E3|55=uatom|54=2|44=2916090|59=1|

# Contributing
1. Fork the Project

2. Create your Feature Branch (git checkout -b feature/NewFeature)

3. Commit your Changes (git commit -m 'Add some NewFeature')

4. Push to the Branch (git push origin feature/NewFeature)

5. Open a Pull Request


# License
Distributed under the Apache License 2.0. See [LICENSE](https://github.com/jim380/re_client/blob/main/LICENSE) for more information.




