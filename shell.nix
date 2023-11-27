{ pkgs ? import <nixpkgs> {}
, golang ? pkgs.go
, godep ? pkgs.dep }:
pkgs.stdenv.mkDerivation {
    name = "go-shell";

    buildInputs = [ golang
                    godep
                  ];

    shellHook = ''
        export GOPATH=`pwd`
        export PATH=$GOPATH/bin:$PATH
    '';
}