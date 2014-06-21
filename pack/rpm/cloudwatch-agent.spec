#
# Spec for bundle cloudwatch-agent into an RPM package
#
Summary: A CloudWatch long running daemon
Name: cloudwatch-agent
Version: 0.0.8
Release: 1
License: MIT
Group: System Environment/Daemons
Source: %{name}.tar.gz
URL: http://github.com/wdalmut/cloudwatch-agent
Distribution: WSS Linux
Vendor: Corley S.r.l.
Packager: Walter Dal Mut <walter.dalmut@corley.it>

%define remote_pack https://github.com/wdalmut/%{name}/archive/%{version}.tar.gz

%description
A long running application metrics collector daemon.
The daemon collects all metrics through UDP/IP socket
and send data collected periodically to AWS CloudWatch

%prep
wget -O %{_sourcedir}/%{name}.tar.gz %{remote_pack}
rm -rf %{_builddir}/%{name}
mkdir -p %{_builddir}/%{name}
zcat %{_sourcedir}/%{name}.tar.gz | tar -xvf -

%build
mv %{name}-%{version}/* %{name}/
cd %{name}
go build -a

%install
mkdir -p %{buildroot}%{_sbindir}
mkdir -p %{buildroot}%{_initrddir}
mkdir -p %{buildroot}%{_sysconfdir}/default
mkdir -p %{buildroot}%{_sysconfdir}/%{name}
cp %{name}/%{name} %{buildroot}%{_sbindir}/
cp %{name}/pack/scripts/cw-agent %{buildroot}%{_initrddir}/
chmod a+x %{buildroot}/%{_initrddir}/cw-agent
cp %{name}/pack/scripts/%{name}.default %{buildroot}%{_sysconfdir}/default/%{name}
cp %{name}/pack/config/%{name}.conf %{buildroot}%{_sysconfdir}/%{name}/%{name}.conf

%files
%doc %{name}/README.md
%{_sbindir}/*
%{_initrddir}/*
%{_sysconfdir}/default/*
%{_sysconfdir}/%{name}/*

%clean
rm -rf %{buildroot}

