export namespace main {
	
	export class BotInfo {
	    name: string;
	    iconPath: string;
	    tomlPath: string;
	
	    static createFrom(source: any = {}) {
	        return new BotInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.iconPath = source["iconPath"];
	        this.tomlPath = source["tomlPath"];
	    }
	}

}

