Vue.component('damage', {
	props: {
	},
	template: `<div class="metric-panel" aria-hidden="true"></div>`,
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
			for (let d = 0; d < this.damages.length; d++) {
				const item = this.damages[d];
				this.fillProcess(this, item);
			}
			this.$bus.$emit('display', this.pies);
		},
		lifeDamage() {
			this.active = 'damage/life';
			this.initProcess(this, 'Damage', 0, 1);
			this.initLives(this);
			for (let d = 0; d < this.damages.length; d++) {
				const item = this.damages[d];
				while (this.killIndex < this.kills.length && this.kills[this.killIndex].time < item.time) {
					this.negatePartial(this, this.kills[this.killIndex]);
					this.killIndex++;
				}
				this.fillPartial(this, item);
			}
			this.fillTotals(this);
			this.$bus.$emit('display', this.pies);
		},
		typeDamage() {
			this.active = 'damage/type';
			const types = ['emp', 'thermal', 'kinetic'];
			this.initCustom(this, 'Types', types, 0, 1);
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
			this.$bus.$emit('display', this.pies);
		}
	},
	mounted() {
		this._boundDamage = (game, damages, heals, kills) => {
			this.game = game;
			this.damages = damages;
			this.kills = kills;
			this.playerDamage();
		};
		this.$bus.$on('damage', this._boundDamage);
	},
	beforeDestroy() {
		this.$bus.$off('damage', this._boundDamage);
	}
});
