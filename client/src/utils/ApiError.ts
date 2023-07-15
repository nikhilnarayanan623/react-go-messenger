class ApiError extends Error {
    private data: any;
  
    constructor(message: string, customData: any) {
      super(message);
      this.name = this.constructor.name;
      this.data = customData;
  
      Object.setPrototypeOf(this,ApiError.prototype);
    }
  
    getCustomData(): any {
      return this.data;
    }
  }
  
  export default ApiError;
  