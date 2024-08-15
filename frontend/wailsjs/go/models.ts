export namespace main {
	
	export class BotDetails {
	    description: string;
	    funFact: string;
	    sourceLink: string;
	    developer: string;
	    language: string;
	    tags: string[];
	
	    static createFrom(source: any = {}) {
	        return new BotDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.description = source["description"];
	        this.funFact = source["funFact"];
	        this.sourceLink = source["sourceLink"];
	        this.developer = source["developer"];
	        this.language = source["language"];
	        this.tags = source["tags"];
	    }
	}
	export class BotSettings {
	    name: string;
	    looksConfig: string;
	    location: string;
	    logoFile: string;
	    runCommand: string;
	    runCommandLinux: string;
	    hivemind: boolean;
	
	    static createFrom(source: any = {}) {
	        return new BotSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.looksConfig = source["looksConfig"];
	        this.location = source["location"];
	        this.logoFile = source["logoFile"];
	        this.runCommand = source["runCommand"];
	        this.runCommandLinux = source["runCommandLinux"];
	        this.hivemind = source["hivemind"];
	    }
	}
	export class BotConfig {
	    settings: BotSettings;
	    details: BotDetails;
	
	    static createFrom(source: any = {}) {
	        return new BotConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.settings = this.convertValues(source["settings"], BotSettings);
	        this.details = this.convertValues(source["details"], BotDetails);
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
	
	export class BotInfo {
	    config: BotConfig;
	    tomlPath: string;
	
	    static createFrom(source: any = {}) {
	        return new BotInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.config = this.convertValues(source["config"], BotConfig);
	        this.tomlPath = source["tomlPath"];
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
	
	export class Result {
	    success: boolean;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	    }
	}
	export class StartMatchOptions {
	    map: string;
	    gameMode: string;
	    blueBots: BotInfo[];
	    orangeBots: BotInfo[];
	
	    static createFrom(source: any = {}) {
	        return new StartMatchOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.map = source["map"];
	        this.gameMode = source["gameMode"];
	        this.blueBots = this.convertValues(source["blueBots"], BotInfo);
	        this.orangeBots = this.convertValues(source["orangeBots"], BotInfo);
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

