Vue.component('display', {
	props: {
	},
	template: `<div class="display row">

		<div class="display" v-for="(pie, id) in pies" :key="'pie-' + id">
			<div class="title-box team-heading">
				<p class="title" style="font-size:22px;margin:4px 0;">{{ teamHeading(id) }}</p>
			</div>
			<display-pie class="display-box" :index="id" :data="pie"></display-pie>
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
				return 'Team ' + (id + 1) + ': ' + this.teams[id].join(', ');
			}
			return 'Team ' + (id + 1);
		},
		loadData(pies) {
			this.pies = pies;
			console.log('display data are ready');
		},
		setTeams(teams) {
			this.teams = teams;
		},
	},
	computed: {

	},
	mounted() {
		this.$root.$on('display', (pies) => {
			this.loadData(pies);
		});
		this.$root.$on('display-teams', (teams) => {
			this.setTeams(teams);
		});
	}
});
