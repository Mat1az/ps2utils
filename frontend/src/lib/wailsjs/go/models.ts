export namespace models {
	
	export class ZSOHeader {
	    is_zso: boolean;
	    magic: string;
	    header_size: number;
	    orig_size: string;
	    block_size: number;
	    version: number;
	    index_shift: number;
	    unused: string;
	
	    static createFrom(source: any = {}) {
	        return new ZSOHeader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.is_zso = source["is_zso"];
	        this.magic = source["magic"];
	        this.header_size = source["header_size"];
	        this.orig_size = source["orig_size"];
	        this.block_size = source["block_size"];
	        this.version = source["version"];
	        this.index_shift = source["index_shift"];
	        this.unused = source["unused"];
	    }
	}
	export class Game {
	    format: number;
	    path: string;
	    id: string;
	    name: string;
	    size: string;
	    opl: boolean;
	    hdl: boolean;
	    zso: ZSOHeader;
	
	    static createFrom(source: any = {}) {
	        return new Game(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.format = source["format"];
	        this.path = source["path"];
	        this.id = source["id"];
	        this.name = source["name"];
	        this.size = source["size"];
	        this.opl = source["opl"];
	        this.hdl = source["hdl"];
	        this.zso = this.convertValues(source["zso"], ZSOHeader);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

