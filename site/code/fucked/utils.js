Vue.mixin({
    methods: {
        getColor(a) {
            switch (a) {
                case 0:
                    return "#EEE";
                case 1:
                    return "#7D0";
                case 2:
                    return "#FB1";
                case 3:
                    return "#0BC";
                case 4:
                    return "#F31";
                case 5:
                    return "#390";
                case 6:
                    return "#70D";
                case 7:
                    return "#E23";
                case 8:
                    return "#33E";
                case 9:
                    return "#721";
                case 10:
                    return "#046";
                case 11:
                    return "#FA5";
                case 12:
                    return "#E0E";
            }
        },
        extract(b, c, d) {
            const e = b.indexOf(c) + c.length;
            if (d) {
                const f = b.indexOf(d, e);
                return b.substring(e, f);
            } else {
                return b.substring(e);
            }
        },
        extractTrim(g, h, i) {
            return this.extract(g, h, i).trim();
        },
        extractFrom(j, k, l) {
            if (l) {
                const m = j.indexOf(l, k);
                return j.substring(k, m);
            } else {
                return j.substring(k);
            }
        },
        extractNumber(n, o, p) {
            const q = n.indexOf(" ", p);
            return parseInt(n.substring(o, q).trim());
        },
        formatNumber(r) {
            let s = "";
            if (r > 100) {
                s += Math.round(r);
                let t = 3;
                while (s.length > t) {
                    const u = s.length - t;
                    s = s.substring(0, u) + " " + s.substring(u);
                    t += 4;
                }
            } else {
                s += r;
                s = s.substring(0, 4);
            }
            return s;
        },
        formatTime(v) {
            return Math.floor(v / 60) + " min " + Math.round(v % 60) + " sec";
        },
        computeTime(w) {
            const x = w.split(":");
            if (x.length === 3) {
                return parseInt(x[0]) * 3600 + parseInt(x[1]) * 60 + parseInt(x[2]);
            } else if (x.length === 2) {
                return parseInt(x[0]) * 60 + parseInt(x[1]);
            } else {
                return parseInt(x[0]);
            }
        },
        buildPie(name, data) {
            const aa = [{
                name: name,
                amount: 0,
                average: 0
            }];
            for (let ab = 0; ab < data.length; ab++) {
                aa.push({
                    name: data[ab],
                    amount: 0
                });
            }
            return aa;
        },
        buildChart(name, data) {
            const ae = [{
                name: "Average " + name,
                data: [],
                step: []
            }];
            for (let af = 0; af < data.length; af++) {
                ae.push({
                    name: data[af],
                    data: [],
                    step: []
                });
            }
            return ae;
        },
        buildTable(name, ah, ai) {
            const aj = [];
            aj.push(new Array(ai.length + 2));
            for (let ak = 0; ak < ah.length; ak++) {
                const al = new Array(ai.length + 2).fill(0);
                al[0] = ah[ak];
                aj.push(al);
            }
            for (let am = 0; am < ai.length; am++) {
                aj[0][am + 1] = ai[am];
            }
            aj.push(new Array(ai.length + 2).fill(0));
            aj[0][0] = name;
            aj[ah.length + 1][0] = "Total";
            aj[0][ai.length + 1] = "Total";
            return aj;
        },
        initProcess(ctx, ao, ap, aq) {
            const ar = ctx.game.teams;
            console.log(ao + " computing have begun");
            ctx.pies = [];
            ctx.charts = [];
            ctx.tables = [];
            ctx.pies.push(this.buildPie(ao, ar[0]));
            ctx.pies.push(this.buildPie(ao, ar[1]));
            ctx.charts.push(this.buildChart(ao, ar[0]));
            ctx.charts.push(this.buildChart(ao, ar[1]));
            ctx.tables.push(this.buildTable(ao, ar[0], ar[aq]));
            ctx.tables.push(this.buildTable(ao, ar[1], ar[ap]));
        },
        initCustom(ctx, at, au, av, aw) {
            const ax = ctx.game.teams;
            console.log(at + " computing have begun");
            ctx.pies = [];
            ctx.charts = [];
            ctx.tables = [];
            ctx.pies.push(this.buildPie(at, au));
            ctx.pies.push(this.buildPie(at, au));
            ctx.charts.push(this.buildChart(at, au));
            ctx.charts.push(this.buildChart(at, au));
            ctx.tables.push(this.buildTable(at, ax[av], au));
            ctx.tables.push(this.buildTable(at, ax[aw], au));
        },
        initLives(ctx) {
            ctx.killIndex = 0;
            ctx.totalLives = [new Array(ctx.game.teams[0].length).fill(1), new Array(ctx.game.teams[1].length).fill(1)];
            ctx.teamLives = [ctx.game.teams[0].length, ctx.game.teams[1].length];
        },
        fillProcess(ctx, ba) {
            const bb = ctx.game;
            const bc = bb.teams;
            ctx.pies[ba.side][0].amount += ba.amount;
            ctx.pies[ba.side][0].average += ba.amount / bc[ba.side].length;
            ctx.pies[ba.side][ba.source + 1].amount += ba.amount;
            const bd = ctx.charts[ba.side][ba.source + 1].step.length;
            const be = ctx.charts[ba.side][0].step.length;
            ctx.charts[ba.side][0].step[be] = ba.time - bb.init;
            ctx.charts[ba.side][0].data[be] = ctx.pies[ba.side][0].average;
            ctx.charts[ba.side][ba.source + 1].step[bd] = ctx.charts[ba.side][0].step[be];
            ctx.charts[ba.side][ba.source + 1].data[bd] = ctx.pies[ba.side][ba.source + 1].amount;
            ctx.tables[ba.side][ba.source + 1][ba.target + 1] += ba.amount;
            ctx.tables[ba.side][bc[ba.side].length + 1][ba.target + 1] += ba.amount;
            ctx.tables[ba.side][ba.source + 1][bc[ba.other].length + 1] += ba.amount;
            ctx.tables[ba.side][bc[ba.side].length + 1][bc[ba.other].length + 1] += ba.amount;
        },
        fillCustom(ctx, bg, bh) {
            const game = ctx.game;
            const teams = game.teams;
            ctx.pies[bg.side][0].amount += bg.amount;
            ctx.pies[bg.side][0].average += bg.amount / bh.length;
            ctx.pies[bg.side][bg.type + 1].amount += bg.amount;
            const bk = ctx.charts[bg.side][bg.type + 1].step.length;
            const bl = ctx.charts[bg.side][0].step.length;
            ctx.charts[bg.side][0].step[bl] = bg.time - game.init;
            ctx.charts[bg.side][0].data[bl] = ctx.pies[bg.side][0].average;
            ctx.charts[bg.side][bg.type + 1].step[bk] = ctx.charts[bg.side][0].step[bl];
            ctx.charts[bg.side][bg.type + 1].data[bk] = ctx.pies[bg.side][bg.type + 1].amount;
            ctx.tables[bg.side][bg.player + 1][bg.type + 1] += bg.amount;
            ctx.tables[bg.side][teams[bg.side].length + 1][bg.type + 1] += bg.amount;
            ctx.tables[bg.side][bg.player + 1][bh.length + 1] += bg.amount;
            ctx.tables[bg.side][teams[bg.side].length + 1][bh.length + 1] += bg.amount;
        },
        negatePartial(ctx, bn) {
            const game = ctx.game;
            const teams = game.teams;
            ctx.totalLives[bn.other][bn.target]++;
            ctx.teamLives[bn.other]++;
            const bq = (ctx.totalLives[bn.other][bn.target] - 1) / ctx.totalLives[bn.other][bn.target];
            const br = (ctx.teamLives[bn.other] - 1) / ctx.teamLives[bn.other];
            ctx.pies[bn.other][0].average *= br;
            ctx.pies[bn.other][bn.target + 1].amount *= bq;
            const bs = ctx.charts[bn.other][bn.target + 1].step.length;
            const bt = ctx.charts[bn.other][0].step.length;
            ctx.charts[bn.other][0].step[bt] = bn.time - ctx.game.init;
            ctx.charts[bn.other][0].data[bt] = ctx.pies[bn.other][0].average;
            ctx.charts[bn.other][bn.target + 1].step[bs] = ctx.charts[bn.other][0].step[bt];
            ctx.charts[bn.other][bn.target + 1].data[bs] = ctx.pies[bn.other][bn.target + 1].amount;
            for (let bu = 0; bu < ctx.game.teams[bn.side].length; bu++) {
                ctx.tables[bn.other][bn.target + 1][bu + 1] *= bq;
            }
            ctx.tables[bn.other][bn.target + 1][teams[bn.side].length + 1] = ctx.pies[bn.other][bn.target + 1].amount;
        },
        fillPartial(ctx, bw) {
            const game = ctx.game;
            const teams = game.teams;
            ctx.pies[bw.side][0].amount += bw.amount;
            ctx.pies[bw.side][0].average += bw.amount / ctx.teamLives[bw.side];
            ctx.pies[bw.side][bw.source + 1].amount += bw.amount / ctx.totalLives[bw.side][bw.source];
            const bz = ctx.charts[bw.side][bw.source + 1].step.length;
            const ca = ctx.charts[bw.side][0].step.length;
            ctx.charts[bw.side][0].step[ca] = bw.time - ctx.game.init;
            ctx.charts[bw.side][0].data[ca] = ctx.pies[bw.side][0].average;
            ctx.charts[bw.side][bw.source + 1].step[bz] = ctx.charts[bw.side][0].step[ca];
            ctx.charts[bw.side][bw.source + 1].data[bz] = ctx.pies[bw.side][bw.source + 1].amount;
            ctx.tables[bw.side][bw.source + 1][bw.target + 1] += bw.amount / ctx.totalLives[bw.side][bw.source];
            ctx.tables[bw.side][bw.source + 1][teams[bw.other].length + 1] = ctx.pies[bw.side][bw.source + 1].amount;
        },
        fillTotals(ctx) {
            const game = ctx.game;
            const totals = [0, 0];
            for (let ce = 0; ce < 2; ce++) {
                for (let cf = 0; cf < game.teams[ce].length; cf++) {
                    totals[ce] += ctx.pies[ce][cf + 1].amount;
                }
                ctx.pies[ce][0].amount = totals[ce];
                ctx.tables[ce][game.teams[ce].length + 1][game.teams[ce > 0 ? 0 : 1].length + 1] = totals[ce];
            }
        }
    }
});