from setuptools import setup, find_packages

setup(
	name="WhatsDifferent",
	version="1.2.0",
	packages=find_packages(),
	description="A tool for analyzing file differences with detailed change tracking",
	author="OG-Open-Source",
	author_email="no-reply@ogtt.tk",
	url="https://github.com/OG-Open-Source/WhatsDifferent",
	classifiers=[
		"Programming Language :: Python :: 3",
		"License :: OSI Approved :: MIT License",
		"Operating System :: OS Independent",
	],
	python_requires=">=3.6",
)