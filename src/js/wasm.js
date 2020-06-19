const wasmBrowserInstantiate = async (wasmModuleUrl, importObject) => {
    let response = undefined;
  
    // Check if the browser supports streaming instantiation
    if (WebAssembly.instantiateStreaming) {
      // Fetch the module, and instantiate it as it is downloading
      response = await WebAssembly.instantiateStreaming(
        fetch(wasmModuleUrl),
        importObject
      );
    } else {
      // Fallback to using fetch to download the entire module
      // And then instantiate the module
      const fetchAndInstantiateTask = async () => {
        const wasmArrayBuffer = await fetch(wasmModuleUrl).then(response =>
          response.arrayBuffer()
        );
        return WebAssembly.instantiate(wasmArrayBuffer, importObject);
      };
  
      response = await fetchAndInstantiateTask();
    }
  
    return response;
  };

  const runWasmApp = async () => {
    const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

    // Get the importObject from the go instance.
    const importObject = go.importObject;
  
    // Instantiate our wasm module
    const wasmModule = await wasmBrowserInstantiate("./asm/sudoku.wasm", importObject);
  
    // Allow the wasm_exec go instance, bootstrap and execute our wasm module
    go.run(wasmModule.instance);
  };
  runWasmApp();