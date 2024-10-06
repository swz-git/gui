export namespace flat {
	
	export class ColorT {
	    a: number;
	    r: number;
	    g: number;
	    b: number;
	
	    static createFrom(source: any = {}) {
	        return new ColorT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.a = source["a"];
	        this.r = source["r"];
	        this.g = source["g"];
	        this.b = source["b"];
	    }
	}
	export class LoadoutPaintT {
	    car_paint_id: number;
	    decal_paint_id: number;
	    wheels_paint_id: number;
	    boost_paint_id: number;
	    antenna_paint_id: number;
	    hat_paint_id: number;
	    trails_paint_id: number;
	    goal_explosion_paint_id: number;
	
	    static createFrom(source: any = {}) {
	        return new LoadoutPaintT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.car_paint_id = source["car_paint_id"];
	        this.decal_paint_id = source["decal_paint_id"];
	        this.wheels_paint_id = source["wheels_paint_id"];
	        this.boost_paint_id = source["boost_paint_id"];
	        this.antenna_paint_id = source["antenna_paint_id"];
	        this.hat_paint_id = source["hat_paint_id"];
	        this.trails_paint_id = source["trails_paint_id"];
	        this.goal_explosion_paint_id = source["goal_explosion_paint_id"];
	    }
	}
	export class MutatorSettingsT {
	    match_length: number;
	    max_score: number;
	    multi_ball: number;
	    overtime_option: number;
	    series_length_option: number;
	    game_speed_option: number;
	    ball_max_speed_option: number;
	    ball_type_option: number;
	    ball_weight_option: number;
	    ball_size_option: number;
	    ball_bounciness_option: number;
	    boost_option: number;
	    rumble_option: number;
	    boost_strength_option: number;
	    gravity_option: number;
	    demolish_option: number;
	    respawn_time_option: number;
	    max_time_option: number;
	    game_event_option: number;
	    audio_option: number;
	
	    static createFrom(source: any = {}) {
	        return new MutatorSettingsT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.match_length = source["match_length"];
	        this.max_score = source["max_score"];
	        this.multi_ball = source["multi_ball"];
	        this.overtime_option = source["overtime_option"];
	        this.series_length_option = source["series_length_option"];
	        this.game_speed_option = source["game_speed_option"];
	        this.ball_max_speed_option = source["ball_max_speed_option"];
	        this.ball_type_option = source["ball_type_option"];
	        this.ball_weight_option = source["ball_weight_option"];
	        this.ball_size_option = source["ball_size_option"];
	        this.ball_bounciness_option = source["ball_bounciness_option"];
	        this.boost_option = source["boost_option"];
	        this.rumble_option = source["rumble_option"];
	        this.boost_strength_option = source["boost_strength_option"];
	        this.gravity_option = source["gravity_option"];
	        this.demolish_option = source["demolish_option"];
	        this.respawn_time_option = source["respawn_time_option"];
	        this.max_time_option = source["max_time_option"];
	        this.game_event_option = source["game_event_option"];
	        this.audio_option = source["audio_option"];
	    }
	}
	export class PlayerLoadoutT {
	    team_color_id: number;
	    custom_color_id: number;
	    car_id: number;
	    decal_id: number;
	    wheels_id: number;
	    boost_id: number;
	    antenna_id: number;
	    hat_id: number;
	    paint_finish_id: number;
	    custom_finish_id: number;
	    engine_audio_id: number;
	    trails_id: number;
	    goal_explosion_id: number;
	    loadout_paint?: LoadoutPaintT;
	    primary_color_lookup?: ColorT;
	    secondary_color_lookup?: ColorT;
	
	    static createFrom(source: any = {}) {
	        return new PlayerLoadoutT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.team_color_id = source["team_color_id"];
	        this.custom_color_id = source["custom_color_id"];
	        this.car_id = source["car_id"];
	        this.decal_id = source["decal_id"];
	        this.wheels_id = source["wheels_id"];
	        this.boost_id = source["boost_id"];
	        this.antenna_id = source["antenna_id"];
	        this.hat_id = source["hat_id"];
	        this.paint_finish_id = source["paint_finish_id"];
	        this.custom_finish_id = source["custom_finish_id"];
	        this.engine_audio_id = source["engine_audio_id"];
	        this.trails_id = source["trails_id"];
	        this.goal_explosion_id = source["goal_explosion_id"];
	        this.loadout_paint = this.convertValues(source["loadout_paint"], LoadoutPaintT);
	        this.primary_color_lookup = this.convertValues(source["primary_color_lookup"], ColorT);
	        this.secondary_color_lookup = this.convertValues(source["secondary_color_lookup"], ColorT);
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
	export class PlayerClassT {
	
	
	    static createFrom(source: any = {}) {
	        return new PlayerClassT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class PlayerConfigurationT {
	    // Go type: PlayerClassT
	    variety?: any;
	    name: string;
	    team: number;
	    location: string;
	    run_command: string;
	    loadout?: PlayerLoadoutT;
	    spawn_id: number;
	    hivemind: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PlayerConfigurationT(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.variety = this.convertValues(source["variety"], null);
	        this.name = source["name"];
	        this.team = source["team"];
	        this.location = source["location"];
	        this.run_command = source["run_command"];
	        this.loadout = this.convertValues(source["loadout"], PlayerLoadoutT);
	        this.spawn_id = source["spawn_id"];
	        this.hivemind = source["hivemind"];
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
	    bluePlayers: any[];
	    orangePlayers: any[];
	    mutatorSettings: flat.MutatorSettingsT;
	
	    static createFrom(source: any = {}) {
	        return new StartMatchOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.map = source["map"];
	        this.gameMode = source["gameMode"];
	        this.bluePlayers = source["bluePlayers"];
	        this.orangePlayers = source["orangePlayers"];
	        this.mutatorSettings = this.convertValues(source["mutatorSettings"], flat.MutatorSettingsT);
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

