Vue.component('heal', {
	props: {
	},
	template: `<div class="metric-panel" aria-hidden="true"></div>`,
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
			for (let h = 0; h < this.heals.length; h++) {
				const item = this.heals[h];
				this.fillProcess(this, item);
			}
			this.$bus.$emit('display', this.pies);
		},
		lifeHeal() {
			this.active = 'heal/life';
			this.initProcess(this, 'Heal', 1, 0);
			this.initLives(this);
			for (let d = 0; d < this.heals.length; d++) {
				const item = this.heals[d];
				while (this.killIndex < this.kills.length && this.kills[this.killIndex].time < item.time) {
					this.negatePartial(this, this.kills[this.killIndex]);
					this.killIndex++;
				}
				this.fillPartial(this, item);
			}
			this.fillTotals(this);
			this.$bus.$emit('display', this.pies);
		},
		typeHeal() {
			this.active = 'heal/type';
			const types = ['shield', 'armor'];
			this.initCustom(this, 'Types', types, 0, 1);
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
			this.$bus.$emit('display', this.pies);
		}
	},
	mounted() {
		this._boundHeal = (game, damages, heals, kills) => {
			this.game = game;
			this.heals = heals;
			this.kills = kills;
			this.playerHeal();
		};
		this.$bus.$on('heal', this._boundHeal);
	},
	beforeDestroy() {
		this.$bus.$off('heal', this._boundHeal);
	}
});
