// chroma.js - TinyGo WASM runtime initialization for Chroma syntax highlighter

class ChromaWASM {
  constructor() {
    this.go = null;
    this.wasm = null;
    this.ready = false;
    this.readyPromise = this.init();
  }

  async init() {
    try {
      // Create a new Go instance (wasm_exec.js already imported in initChroma)
      const go = new Go();
      WebAssembly.instantiateStreaming(fetch("./static/chroma.wasm"), go.importObject).then((result) => {
          go.run(result.instance);
          this.ready = true;
      });
      console.log("Chroma WASM module initialized successfully");
    } catch (error) {
      console.error("Failed to initialize Chroma WASM module:", error);
      throw error;
    }
  }

  async waitForReady() {
    await this.readyPromise;
    if (!this.ready) {
      throw new Error("Chroma WASM module failed to initialize");
    }
  }

  async highlight(source, lexer, formatter, withClasses) {
    await this.waitForReady();

    if (typeof window.highlight !== "function") {
      throw new Error("highlight function not available from WASM module");
    }

    try {
      return window.highlight(source, lexer, formatter, withClasses);
    } catch (error) {
      console.error("Error calling highlight function:", error);
      throw error;
    }
  }
}

export function isWasmSupported() {
  try {
    if (
      typeof WebAssembly === "object" &&
      typeof WebAssembly.instantiate === "function"
    ) {
      // The smallest possible WebAssembly module (magic number + version)
      const module = new WebAssembly.Module(
        Uint8Array.of(0x0, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00),
      );
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

async function initChroma() {
  if (!isWasmSupported()) {
    return null;
  }

  try {
    await import("./wasm_exec.js");
    return new ChromaWASM();
  } catch (error) {
    return null;
  }
}

// Create global instance, null if WASM is not supported or file doesn't exist.
export const chroma = await initChroma();
