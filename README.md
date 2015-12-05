![Build status](https://circleci.com/gh/sebdah/lion.png?circle-token=ed5f993b299f1174bb8a7aa960fe154b409873a0 "Build status")

# Introduction

Lion is a basic localization service using `git` as the backend for storing
translations.

# Git repository structure

The following structure is expected in the localization repositories.

	<project-name>/
		source/
			<resource>.json
		<language>_<region>/
			<resource>.json
