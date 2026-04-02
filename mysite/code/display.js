Vue.component('display', {
	props: {
	},
	template: `<div class="charts-section">
		<div class="title-box charts-section-title">
			<p class="title" style="font-size:28px;margin:8px 0;">Графики (две команды)</p>
		</div>
		<div class="display row">

			<div class="display chart-column" v-for="(pie, id) in pies" :key="'pie-' + id">
				<div class="title-box team-heading">
					<p class="title team-title">{{ teamHeading(id) }}</p>
				</div>
				<display-pie class="display-box" :index="id" :data="pie"></display-pie>
			</div>

		</div>
	</div>`,
	data() {
		return {
			pies: [[], []],
			teams: null,
		}
	},
	methods: {
		teamHeading(id) {
			if (this.teams && this.teams[id] && this.teams[id].length) {
				return 'Команда ' + (id + 1) + ': ' + this.teams[id].join(', ');
			}
			return 'Команда ' + (id + 1);
		},
		loadData(pies) {
			if (!pies || !Array.isArray(pies)) {
				this.pies = [[], []];
			} else {
				const a = pies[0];
				const b = pies[1];
				this.pies = [
					Array.isArray(a) ? a.slice() : [],
					Array.isArray(b) ? b.slice() : [],
				];
			}
			console.log('display: pies updated', this.pies);
		},
		setTeams(teams) {
			this.teams = teams;
		},
	},
	mounted() {
		const bus = this.$bus;
		this._boundDisplay = (pies) => this.loadData(pies);
		this._boundTeams = (teams) => this.setTeams(teams);
		bus.$on('display', this._boundDisplay);
		bus.$on('display-teams', this._boundTeams);
	},
	beforeDestroy() {
		const bus = this.$bus;
		bus.$off('display', this._boundDisplay);
		bus.$off('display-teams', this._boundTeams);
	}
});
