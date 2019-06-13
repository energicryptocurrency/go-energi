module.exports = {
    "env": {
        "commonjs": true,
        "es6": true,
        "mocha": true,
        "node": true,
    },
    "extends": "eslint:recommended",
    "globals": {
        "Atomics": "readonly",
        "SharedArrayBuffer": "readonly",
        "expect": "readonly",
        "assert": "readonly",
        "web3": "readonly",
        "artifacts": "readonly",
        "contract": "readonly",
    },
    "parserOptions": {
        "ecmaVersion": 2018
    },
    "rules": {
        "indent": [
            "error",
            4
        ],
        "linebreak-style": [
            "error",
            "unix"
        ],
        "no-unused-vars": [
            "error",
            {
                "vars" : "all",
                "args" : "all",
                "argsIgnorePattern" : "^(_.*)$"
            }
        ],
    }
};
