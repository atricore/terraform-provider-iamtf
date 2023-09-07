# JOSSO development environment: tools to work with JOSSO/IAM.tf

{
  description = "Terraform provider IAM.tf/JOSSO dev shell";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/release-23.05";
    flake-utils.url = "github:numtide/flake-utils";
    josso-pkgs.url = "github:sgonzalezoyuela/mynix-pkgs";
  };
  outputs = { self, nixpkgs, flake-utils, josso-pkgs }:
    flake-utils.lib.eachDefaultSystem
      (system:
        let

          pkgs = import nixpkgs {
            inherit  system;
          };
 

        in
        with pkgs;
        {
          devShells.default = mkShell {
            buildInputs = [
              gnumake
              go
              gotools
              gopls
              go-outline
              gocode
              gopkgs
              gocode-gomod
              godef
              golint
              gofumpt
              terraform
            ];
            shellHook = ''
             echo "Terraform provider JOSSO/IAM.tf GO environment: ${go.name} ${terraform.name}"
            '';
          };
        }
      );
}
