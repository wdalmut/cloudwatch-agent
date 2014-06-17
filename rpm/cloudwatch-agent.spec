#
# Spec for bundle cloudwatch-agent into an RPM package
#
Summary: A CloudWatch long running daemon
Name: cloudwatch-agent
Version: 0.0.1
Release: 1
License: MIT
Group: System Environment/Daemons
Source: %{name}.tar.gz
URL: http://github.com/wdalmut/cloudwatch-agent
Distribution: WSS Linux
Vendor: Corley S.r.l.
Packager: Walter Dal Mut <walter.dalmut@corley.it>

%define remote_pack http://github.com/wdalmut/%{name}/%{version}.tar.gz

%description
A long running application metrics collector daemon.
The daemon collects all metrics through UDP/IP socket
and send data collected periodically to AWS CloudWatch

%prep
wget -O %{name}.tar.gz https://github.com/wdalmut/cloudwatch-agent/archive/%{version}.tar.gz
rm -rf %{_builddir}/%{name}
mkdir -p %{_builddir}/%{name}
zcat %{_sourcedir}/%{name}.tar.gz | tar -xvf -

%build
cd %{name}
go build

%install
mkdir -p %{buildroot}%{_sbindir}
cp %{name}/%{name} %{buildroot}%{_sbindir}/

%files
%doc %{name}/README.md
%{_sbindir}/*

%clean
rm -rf %{buildroot}

