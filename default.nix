{ pkgs ? import <nixpkgs> { }, stdenv ? pkgs.stdenv, lib ? stdenv.lib }:

with pkgs;
with builtins;

buildGoModule rec {
  name = "s3-edit-${version}";
  version = "0.0.13";

  modSha256 = "0naaga35bxah63smis5wlbhyc2lfjjgbnw3i7ipsvk937y80qa5l";

  src = fetchFromGitHub{
    owner = "tsub";
    repo = "s3-edit";
    rev = "v${version}";
    sha256 = "054n67cknhf411z4i6y2rbvkf12rlidwz969sdm68d7ncpi2bjqw";
  };

  goPackagePath = [ "github.com/tsub/s3-edit" ];

  meta = with lib; {
    description = "Edit directly a file on Amazon S3 in CLI";
    homepage = https://github.com/tsub/s3-edit;
    license = licenses.mit;
    platforms = platforms.linux ++ platforms.darwin;
  };
}
