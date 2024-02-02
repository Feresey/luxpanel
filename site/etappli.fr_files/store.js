Vue.component('store', {
	props: {
	},
	template: `<div class="store row">

		<div class="half-row">
			<button class="button256" v-for="(menu, id) in menus" @click="openMenu(menu)" v-bind:class="{ on: active === menu }">{{menu}}</button>
		</div>
		
    	<damage v-show="active === 'damage'"></damage>
    	<heal   v-show="active === 'heal'"  ></heal>
    	<kill   v-show="active === 'kill'"  ></kill>
	
	</div>`,
	data() {
		return {
			active: 'damage',
			menus: ['damage', 'heal', 'kill'],
			game: null,
			damages: [{
				source: 'Eta',
				target: 'ECM',
				amount: 1337,
				time: 512345,
				weapon: 'IDK',
				damage: 'THERMAL',
				critical: false
			}],
			selfs: [],
			heals: [],
			kills: [],
			ignores: ['n/a', 'HealBot', 'WarpGate', 'Debris']
		}
	},
	methods: {
		alertPlayerError(player, line) {
			for (let i=0; i<this.ignores.length; i++) {
				if (player.indexOf(this.ignores[i]) > -1) {
					return;
				}
			}
			console.log('player '+player+' not present in game teams ! time code : '+line.substring(0,8));
		},
		setIndexes(item, teams, line) {
			const indexSourceFriendly = teams[0].indexOf(item.source);
			const indexSourceHostile  = teams[1].indexOf(item.source);
			const indexTargetFriendly = teams[0].indexOf(item.target);
			const indexTargetHostile  = teams[1].indexOf(item.target);
			if (indexSourceFriendly === -1 && indexSourceHostile === -1) {
				this.alertPlayerError(item.source, line);
				return false;
			}
			if (indexTargetFriendly === -1 && indexTargetHostile === -1 ){
				this.alertPlayerError(item.target, line);
				return false;
			}
			if (indexSourceFriendly > -1) {
				item.side = 0;
				item.source = indexSourceFriendly;
			} else {
				item.side = 1;
				item.source = indexSourceHostile;
			}
			if (indexTargetFriendly > -1) {
				item.other = 0;
				item.target = indexTargetFriendly;
			} else {
				item.other = 1;
				item.target = indexTargetHostile;
			}
			return true;
		},
		getDamage(line) {
			if (line.indexOf('THERMAL') > -1) {
				return 1;//'thermal';
			} else if (line.indexOf('KINETIC') > -1) {
				return 2;//'kinetic';
			} else if (line.indexOf('EMP') > -1) {
				return 0;//'emp';
			}
			return -1;//'unknown';
		},
		getHeal(line) {
			if (line.indexOf('ShieldEmitter') > -1 ||   //engi aura
				line.indexOf('FrigateDrone') > -1 ||    //engi drone
				line.indexOf('ShieldRay') > -1 ||       //engi station
				line.indexOf('ShieldRestore') > -1 ){   //self heal
				//missing arena bonus
				return 0;//'shield';
			} else if (line.indexOf('RemoteRepairDrones') > -1 ||   //engi aura
				line.indexOf('Nanobots_Cloud') > -1 ||              //engi weapon
				line.indexOf('RepairRay') > -1 ||                   //engi station
				line.indexOf('RepairDrones') > -1 ||                //self heal
				line.indexOf('HealHull') > -1 ){                    //arena bonus
				return 1;//'armor';
			}
			return -1;//'unknown';
		},
		storeLogs(game, logs) {
			console.log('storing logs have begun');
			this.game = game;
			this.damages = [];
			this.selfs = [];
			this.heals = [];
			this.kills = [];
			for (let l=0; l<logs.length; l++) {
				const line = logs[l];
				if (line.indexOf('| Damage') > -1) {
					const damage = {
						source: this.extractTrim(line, '| Damage', '|'),
						target: this.extractTrim(line, '->', '|'),
						amount: this.extractNumber(line, 88, 92),
						weapon: this.extract(line, ') ', ' '),
						damage: this.getDamage(line),
						time: this.computeTime(line.substring(0,8)),
						critical: (line.indexOf('|CRIT') > -1)
					};
					const valid = this.setIndexes(damage, game.teams, line);
					if (valid) {
						if (damage.side === damage.other) {
							this.selfs.push(damage);
						} else {
							this.damages.push(damage);
						}
					}
				}
				if (line.indexOf('| Heal') > -1) {
					const heal = {
						source: this.extractTrim(line, '| Heal', '|'),
						target: this.extractTrim(line, '->', '|'),
						amount: this.extractNumber(line, 88, 92),
						weapon: this.extract(line, 96),
						heal: this.getHeal(line),
						time: this.computeTime(line.substring(0,8))
					};
					const valid = this.setIndexes(heal, game.teams, line);
					if (valid) {
						if (heal.side === heal.other) {
							if (heal.source === heal.target) {
								heal.self = true;
							}
							this.heals.push(heal);
						}
					}
				}
				if (line.indexOf('| Killed') > -1) {
					const kill = {
						source: this.extract(line, 'killer ', '|'),
						target: this.extract(line, '| Killed ', '\t'),
						amount: 1,
						total: 0,
						participants: [],
						implications: [],
						time: this.computeTime(line.substring(0,8))
					};
					const valid = this.setIndexes(kill, game.teams, line);
					if (valid) {
						if (kill.side !== kill.other) {
							//add participation
							l++;
							while (logs[l].indexOf('|    Participant') > -1) {
								const participation = {
									source: this.extractTrim(logs[l], '|    Participant', 'Ship_'),
									amount: parseInt(this.extract(logs[l], 'totalDamage ', ';'))
								};
								const indexSourceFriendly = game.teams[0].indexOf(participation.source);
								const indexSourceHostile  = game.teams[1].indexOf(participation.source);
								if (indexSourceFriendly === -1 && indexSourceHostile === -1) {
									this.alertPlayerError(participation.source, line);
								} else {
									if (indexSourceFriendly > -1) {
										participation.side = 0;
										participation.source = indexSourceFriendly;
									} else {
										participation.side = 1;
										participation.source = indexSourceHostile;
									}
									if (kill.side === participation.side) {
										kill.total += participation.amount;
										kill.participants.push(participation.source);
										kill.implications.push(participation.amount);
									}
								}
								l++;
							}
							l--;
							//finalize kill
							for (let p=0;p<kill.implications.length;p++) {
								kill.implications[p] /= kill.total;
							}
							this.kills.push(kill);
						}
					}
				}
			}
			this.openMenu(this.active);
		},
		openMenu(type) {
			console.log('menu '+type+' opened');
			this.active = type;
			this.$root.$emit(type, this.game, this.damages, this.heals, this.kills);
		}
	},
	computed: {

	},
	mounted() {
		this.$root.$on('import', (game, logs) => {
			this.storeLogs(game, logs);
		});
	}
});