Vue.component('damage', {
	props: {
	},
	template: `<div class="half-row">

		<button class="button256" @click="playerDamage()"  v-bind:class="{ on: active === 'damage/player' }">Damage / Player</button>
		
		<button class="button256" @click="lifeDamage()"  v-bind:class="{ on: active === 'damage/life' }">Damage / Life</button>
		
		<button class="button256" @click="typeDamage()"  v-bind:class="{ on: active === 'damage/type' }">Damage / Type</button>
	
	</div>`,
	data() {
		return {
			active: 'damage/player',
			pies: [[], []],
			charts: [[], []],
			tables: [[], []],
		}
	},
	methods: {
		playerDamage() {
			this.active = 'damage/player';
			this.initProcess(this, 'Damage', 0, 1);
			//fill damages
			for (let d = 0; d < this.damages.length; d++) {
				const item = this.damages[d];
				this.fillProcess(this, item);
			}
			this.$root.$emit('display', this.pies, this.charts, this.tables);
		},
		lifeDamage() {
			this.active = 'damage/life';
			this.initProcess(this, 'Damage', 0, 1);
			this.initLives(this);
			//fill damages
			for (let d = 0; d < this.damages.length; d++) {
				const item = this.damages[d];
				//negate values
				while (this.killIndex < this.kills.length && this.kills[this.killIndex].time < item.time) {
					this.negatePartial(this, this.kills[this.killIndex]);
					this.killIndex++;
				}
				//fill amounts
				this.fillPartial(this, item);
			}
			//fix total amounts
			this.fillTotals(this);
			this.$root.$emit('display', this.pies, this.charts, this.tables);
		},
		typeDamage() {
			this.active = 'damage/type';
			const types = ['emp', 'thermal', 'kinetic'];
			this.initCustom(this, 'Types', types, 0, 1);
			//fill damages
			for (let d = 0; d < this.damages.length; d++) {
				const damage = this.damages[d];
				const item = {
					player: damage.source,
					type: damage.damage,
					amount: damage.amount,
					time: damage.time,
					side: damage.side,
				};
				if (item.type > -1) {
					this.fillCustom(this, item, types);
				}
			}
			this.$root.$emit('display', this.pies, this.charts, this.tables);
		}
	},
	computed: {

	},
	mounted() {
		this.$root.$on('damage', (game, damages, heals, kills) => {
			this.game = game;
			this.damages = damages;
			this.kills = kills;
			this.playerDamage();
		});
	}
});