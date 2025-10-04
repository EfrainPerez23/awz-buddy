{
    description = "AWZ-BUDDY CLI with Go and S3 support";

    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.05"; 
    }

    outputs = { self, nixpkgs }:
        let
            system = "x86_64-linux";
            pkgs = import nixpkgs { inherit system; };
        in
            devShells.${system}.default = pkgs.mkShell {
                builtInputs = [
                    pkgs.go
                    pkgs.git
                    pkgs.make
                ];

                shellHook = ''
                    echo "Welcome to AWZ-BUDDY dev shell!"
                    echo "Go version: $(go version)"
                '';
            };

            packages.${system}.aws-buddy = pkgs.stdenv.mkDerivation {
                pname = "awz-buddy";
                version  = "0.0.1"
                src = "./";
                builtInputs = [ pkgs.g pkgs.make];

                buildPhase = ''
                    make
                '';

                installPhase = ''
                    mkdir -p $out/bin
                    cp dist/* $out/bin/
                '';

                meta = with pkgs.lib; {
                    description = "AWZ-BUDDY CLI is a lightweight CLI tool to help you manage and audit your AWS resources";
                    license = pkgs.lib.licenses.mit;
                }
            }


}