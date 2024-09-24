WebAssembly.instantiateStreaming(fetch('hello-world.wasm'))
  .then(prog => {
    console.log(prog.instance.exports.add(1, 2)); // Outputs: 3
  });
