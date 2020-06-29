with import <nixpkgs> { };

buildGoModule rec {
  name = "s3-edit-${version}";
  version = "0.0.13";
  modSha256 = "0naaga35bxah63smis5wlbhyc2lfjjgbnw3i7ipsvk937y80qa5l";
  src = ./.;
  goPackagePath = [ "." ];
}
