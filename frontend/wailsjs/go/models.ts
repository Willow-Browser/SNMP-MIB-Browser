export namespace oidstorage {
	
	export class Oid {
	    name: string;
	    oid: string;
	    mib: string;
	    syntax: string;
	    access: string;
	    status: string;
	    defVal: string;
	    indexes: string;
	    description: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new Oid(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.oid = source["oid"];
	        this.mib = source["mib"];
	        this.syntax = source["syntax"];
	        this.access = source["access"];
	        this.status = source["status"];
	        this.defVal = source["defVal"];
	        this.indexes = source["indexes"];
	        this.description = source["description"];
	        this.type = source["type"];
	    }
	}

}

