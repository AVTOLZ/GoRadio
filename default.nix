{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoModule {
  name = "GoRadio";
  src = ./.;
  vendorHash = "sha256-LEoWUO/TyLONd8m0D1hki+WvmLO09JyESuH+eHRWoSc=";
}