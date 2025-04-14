{
  description = "QIF Converter";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages = rec {
          default = build;
          build =
            with pkgs;
            buildGoModule {
              pname = "qif-converter";
              version = "0.0.1";
              src = self;
              vendorHash = null;

              meta = {
                description = "A CLI to convert QIF files exported by Homebank into JSON files";
                maintainers = with lib.maintainers; [ ederoyd46 ];
                license = lib.licenses.mit;
              };
            };

          buildDockerImage = pkgs.dockerTools.buildImage {
            name = "qif-converter";
            tag = "latest";
            config = {
              Content = build;
              Cmd = "${build}/bin/product-graph";
            };
          };
        };

        devShells = rec {
          default = devShell;

          devShell = pkgs.mkShell {
            buildInputs =
              with pkgs;
              [
                go
              ]
              ++ pkgs.lib.optionals pkgs.stdenv.isDarwin [ pkgs.darwin.Security ];
          };
        };
      }
    );
}
