# Create RPM packages

In order to create an RPM package just issue

```
$ rpmbuild -ba cloudwatch-agent.spec
```

## Version

The rpm package download the tag from Github repository using the spec
file version.

