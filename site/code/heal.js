Vue.component('heal', {
	props: {
	},
	template: `<div class="half-row">

		<button class="button256" @click="playerHeal()"  v-bind:class="{ on: active === 'heal/player' }">Heal / Player</button>
		
		<button class="button256" @click="lifeHeal()"  v-bind:class="{ on: active === 'heal/life' }">Heal / Life</button>
		
		<button class="button256" @click="typeHeal()"  v-bind:class="{ on: active === 'heal/type' }">Heal / Type</button>
	
	</div>`,
	data() {
		return {
			active: 'heal/player',
			pies: [[], []],
			charts: [[], []],
			tables: [[], []],
		}
	},
	methods: {
		playerHeal() {
			this.active = 'heal/player';
			this.initProcess(this, 'Heal', 1, 0);
			//fill damages
			for (let h = 0; h < this.heals.length; h++) {
				const item = this.heals[h];
				this.fillProcess(this, item);
			}
			//this.pies[0][0].average = this.pies[0][0].amount / (this.pies[0].length - 1);
			//this.pies[1][0].average = this.pies[1][0].amount / (this.pies[1].length - 1);
			this.$root.$emit('display', this.pies, this.charts, this.tables);
		},
		lifeHeal() {
			this.active = 'heal/life';
			this.initProcess(this, 'Heal', 1, 0);
			this.initLives(this);
			//fill damages
			for (let d = 0; d < this.heals.length; d++) {
				const item = this.heals[d];
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
		typeHeal() {
			this.active = 'heal/type';
			const types = ['shield', 'armor'];
			this.initCustom(this, 'Types', types, 0, 1);
			//fill heals
			for (let d = 0; d < this.heals.length; d++) {
				const heal = this.heals[d];
				const item = {
					player: heal.source,
					type: heal.heal,
					amount: heal.amount,
					time: heal.time,
					side: heal.side,
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
		this.$root.$on('heal', (game, damages, heals, kills) => {
			this.game = game;
			this.heals = heals;
			this.kills = kills;
			this.playerHeal();
		});
	}
});
