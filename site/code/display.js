Vue.component('display', {
	props: {
	},
	template: `<div class="display row">

		<div class="display" v-for="(mode, id) in modes">
			<div class="half-row">
				<button class="button192" v-for="(tab) in tabs" @click="displayMode(id, tab)" v-bind:class="{ on: modes[id] === tab }">{{tab}}</button>
			</div>
			
			<display-pie   class="display-box" v-if="modes[id] === 'pie'" :index="id" :data="pies[id]"></display-pie>
			<display-chart class="display-box" v-if="modes[id] === 'chart'" :index="id" :data="charts[id]"></display-chart>
			<display-table class="display-box" v-if="modes[id] === 'table'" :index="id" :data="tables[id]"></display-table>
		</div>
	
	</div>`,
	data() {
		return {
			tabs: ['pie', 'chart', 'table'],
			modes: ['pie', 'pie'],
			pies: [[], []],
			charts: [[], []],
			tables: [[], []],
		}
	},
	methods: {
		loadData(pies, charts, tables) {
			this.pies = pies;
			this.charts = charts;
			this.tables = tables;
			console.log('display data are ready');
			//this.displayMode(0, 'pie');
			//this.displayMode(1, 'pie');
		},
		displayMode(index, mode) {
			this.modes[index] = mode;
			//force update
			this.modes = [this.modes[0], this.modes[1]];
		}
	},
	computed: {

	},
	mounted() {
		this.$root.$on('display', (pies, charts, tables) => {
			this.loadData(pies, charts, tables);
		});
	}
});