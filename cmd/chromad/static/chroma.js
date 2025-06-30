// chroma.js - TinyGo WASM runtime initialization for Chroma syntax highlighter

// Import wasm_exec.js so that it initialises the Go WASM runtime.
import './wasm_exec.js';

class ChromaWASM {
    constructor() {
        this.go = null;
        this.wasm = null;
        this.ready = false;
        this.readyPromise = this.init();
    }

    async init() {
        try {
            // Create a new Go instance
            this.go = new Go();

            // Load the WASM module
            const wasmResponse = await fetch('./static/chroma.wasm');
            if (!wasmResponse.ok) {
                throw new Error(`Failed to fetch chroma.wasm: ${wasmResponse.status}`);
            }

            const wasmBytes = await wasmResponse.arrayBuffer();
            const wasmModule = await WebAssembly.instantiate(wasmBytes, this.go.importObject);

            this.wasm = wasmModule.instance;

            // Run the Go program
            this.go.run(this.wasm);

            this.ready = true;
            console.log('Chroma WASM module initialized successfully');
        } catch (error) {
            console.error('Failed to initialize Chroma WASM module:', error);
            throw error;
        }
    }

    async waitForReady() {
        await this.readyPromise;
        if (!this.ready) {
            throw new Error('Chroma WASM module failed to initialize');
        }
    }

    async highlight(source, lexer, formatter, withClasses) {
        await this.waitForReady();

        if (typeof window.highlight !== 'function') {
            throw new Error('highlight function not available from WASM module');
        }

        try {
            return window.highlight(source, lexer, formatter, withClasses);
        } catch (error) {
            console.error('Error calling highlight function:', error);
            throw error;
        }
    }
}


export function isWasmSupported() {
  try {
    if (typeof WebAssembly === "object" && typeof WebAssembly.instantiate === "function") {
      // The smallest possible WebAssembly module (magic number + version)
      const module = new WebAssembly.Module(Uint8Array.of(0x0, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00));
      if (module instanceof WebAssembly.Module) {
        // Try to instantiate the module to ensure it's truly runnable
        return new WebAssembly.Instance(module) instanceof WebAssembly.Instance;
      }
    }
  } catch (e) {
    // An error occurred (e.g., due to CSP or other restrictions)
  }
  return false;
}


// Create global instance, null if WASM is not supported.
export const chroma = isWasmSupported() ? new ChromaWASM() : null;
