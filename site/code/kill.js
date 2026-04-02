Vue.component('kill', {
	props: {
	},
	template: `<div class="metric-panel"></div>`,
	data() {
		return {
			active: 'kill/player',
			pies: [[], []],
			charts: [[], []],
			tables: [[], []],
		}
	},
	methods: {
		playerKill() {
			this.active = 'kill/player';
			this.initProcess(this, 'Kill', 0, 1);
			for (let k = 0; k < this.kills.length; k++) {
				const kill = this.kills[k];
				this.fillProcess(this, kill);
			}
			this.$root.$emit('display', this.pies);
		},
		playerParticipation() {
			this.active = 'participation/player';
			this.initProcess(this, 'Participation', 0, 1);
			for (let k = 0; k < this.kills.length; k++) {
				const kill = this.kills[k];
				for (let p = 0; p < kill.participants.length; p++) {
					const item = {
						source: kill.participants[p],
						target: kill.target,
						amount: kill.implications[p] * 100,
						time: kill.time,
						side: kill.side,
						other: kill.other
					};
					this.fillProcess(this, item);
				}
			}
			this.$root.$emit('display', this.pies);
		},
		lifeKill() {
			this.active = 'kill/life';
			this.initProcess(this, 'kill / Life', 0, 1);
			this.initLives(this);
			for (let k = 0; k < this.kills.length; k++) {
				const item = this.kills[k];
				this.negatePartial(this, item);
				this.fillPartial(this, item);
			}
			this.fillTotals(this);
			this.$root.$emit('display', this.pies);
		}
	},
	computed: {

	},
	mounted() {
		this.$root.$on('kill', (game, damages, heals, kills) => {
			this.game = game;
			this.kills = kills;
			this.playerKill();
		});
	}
});
