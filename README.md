# Versioner

A command line based tool to read and create [git tags](https://git-scm.com/book/en/v2/Git-Basics-Tagging) that follow [semantic versioning](https://semver.org/).

## Installation

Download the latest `versioner` binary from the releases page and put it somewhere in [your path](http://www.linfo.org/path_env_var.html), e.g.

```sh
mv /path/to/downloaded/versioner /usr/local/versioner
```

## Usage

Run `versioner` to display the tool information, including all available commands.

## Commands

### latest

Run `versioner latest` to read the last version tag on the current git branch.

### bump

Run `versioner bump` to bump the latest **patch** version up by one number. 

The new version will be shown and you will be prompted to confirm before a git tag is created.

To bump the major or minor version, specify it as a third argument, e.g.

```sh
$ versioner bump major
Latest version is 0.2.0
Next version will be v1.0.0
Create new git tag with this version? [y/n]
```

### help

Coming soon...


## Release History

* 0.1.0
	* First release. Added LICENSE and README files.

## Contributing

1. Fork it (<https://github.com/samdjstevens/versioner/fork>)
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request