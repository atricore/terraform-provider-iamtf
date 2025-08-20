{
  description = "JOSSO/IAM.tf API dev shell (go + openapi-generator-cli)";

  # Define the specific nixpkgs commit as an input for openapi-generator-cli 6.2.1
  inputs.prev.url = "github:NixOS/nixpkgs/1732ee9120e43c1df33a33004315741d0173d0b2";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";

  # Define the outputs
  outputs = { self, prev, nixpkgs }:
    let
      # Linux and Mac
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
        pkgs = import nixpkgs { inherit system; };
        prevs = import prev { inherit system; };
      });

    in
    {
      devShells = forEachSupportedSystem ({ pkgs, prevs }: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gotools
            gopls
            go-outline
            gopkgs
            gocode-gomod
            godef
            golint
            gofumpt
            gnumake
            prevs.openapi-generator-cli
          ];
          shellHook = ''
            echo "JOSSO/IAM.tf GO environment: ${pkgs.go.name} / ${prevs.openapi-generator-cli.name}"
          '';
        };
      });
    };
}
