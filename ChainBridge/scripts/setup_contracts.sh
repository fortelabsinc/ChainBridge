#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

CONTRACTS_REPO="https://github.com/fortelabsinc/chainbridge-solidity"
CONTRACTS_TAG="main"
CONTRACTS_BRANCH="main"
CONTRACTS_DIR="./solidity"
DEST_DIR="./bindings"

set -eux

case $TARGET in
	"build")
		git clone $CONTRACTS_REPO $CONTRACTS_DIR
    pushd $CONTRACTS_DIR
    git checkout $CONTRACTS_TAG

    make install-deps
    make bindings

    popd

    mkdir $DEST_DIR
    cp -r $CONTRACTS_DIR/build/bindings/go/* $DEST_DIR
		;;

	"cli-only")
		git clone -b $CONTRACTS_BRANCH $CONTRACTS_REPO $CONTRACTS_DIR
    pushd $CONTRACTS_DIR
    git checkout $CONTRACTS_COMMIT
		;;

esac
