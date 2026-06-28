'use strict';

// A minimal, dependency-free module used as a curated build-from-source sample
// for Stout end-to-end validation. No install/lifecycle scripts, no network.
function greeting() {
  return 'hello from a Stout curated build';
}

module.exports = { greeting };
