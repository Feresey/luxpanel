Vue.component("eta-chart", {
    props: {
        index: {
            type: Number,
            required: true
        },
        curves: {
            type: Array,
            required: true
        }
    },
    template: `<canvas class="full" :ref="'chart-canvas-'+index">Canvas not enabled</canvas>`,
    data() {
        return {
            canvas: {
                context: null
            },
            ctx: null,
            width: 0,
            height: 0,
            mouse: {
                in: true,
                x: 500,
                y: 100
            }
        };
    },
    methods: {
        mouseIn(a) {
            this.mouse.in = true;
            this.mouse.x = a.offsetX;
            this.mouse.y = a.offsetY;
            this.drawChart();
        },
        mouseMove(c) {
            this.mouse.x = c.offsetX;
            this.mouse.y = c.offsetY;
            console.log(this.mouse.x, this.mouse.y);
            this.drawChart();
        },
        mouseOut() {
            this.mouse.in = true;
            this.drawChart();
        },
        getX(d, e, f, g) {
            return g.L + g.P + (d - e) * f;
        },
        getY(h, i, j, k, m) {
            return m - k.B - k.P - (h - i) * j;
        },
        getLastIndex(n, o) {
            let q = -1;
            for (let s = 0; s < n.step.length; s++) {
                if (n.step[s] <= o) {
                    q = s;
                } else {
                    break;
                }
            }
            return q;
        },
        drawChart() {
            this.ctx.fillStyle = "#000";
            this.ctx.fillRect(0, 0, this.width, this.height);
            const curves = this.curves;
            if (!curves || !curves.length) {
                return;
            }
            const width = this.width;
            const height = this.height;
            if (width === 0 || height === 0 || curves.length === 0) {
                return;
            }
            const z = {
                T: 8,
                L: 50,
                B: 20,
                R: 8,
                P: 0
            };
            let ab = -1;
            let ac = -1;
            let ad = null;
            let ae = null;
            let af = null;
            for (let i = 0; i < curves.length; i++) {
                const curve = curves[i];
                if (curve.step.length > 0) {
                    if (ad === null) {
                        ad = curve.step[0] > 999999999;
                        ae = curve.step[0];
                        af = curve.step[curve.step.length - 1];
                    }
                    const ai = curve.step[0];
                    const aj = curve.step[curves[0].step.length - 1];
                    if (ai <= ae && ai !== null) {
                        ae = ai;
                    }
                    if (aj > af && aj !== null) {
                        af = aj;
                    }
                }
                for (let i = 0; i < curve.data.length; i++) {
                    const cdata = curve.data[i];
                    if (ab === -1 || cdata < ab && cdata !== null) {
                        ab = cdata;
                    }
                    if (ac === -1 || cdata > ac && cdata !== null) {
                        ac = cdata;
                    }
                }
            }
            this.ctx.lineWidth = this.width / 200;
            this.ctx.fillStyle = "#EEE";
            this.ctx.strokeStyle = "#EEE";
            this.ctx.beginPath();
            this.ctx.moveTo(z.L, height - z.B + 4);
            this.ctx.lineTo(z.L, z.T);
            this.ctx.lineTo(z.L + 4, z.T + 4);
            this.ctx.lineTo(z.L - 4, z.T + 4);
            this.ctx.lineTo(z.L, z.T);
            this.ctx.moveTo(z.L - 4, height - z.B);
            this.ctx.lineTo(width - z.R, height - z.B);
            this.ctx.lineTo(width - z.R - 4, height - z.B + 4);
            this.ctx.lineTo(width - z.R - 4, height - z.B - 4);
            this.ctx.lineTo(width - z.R, height - z.B);
            this.ctx.stroke();
            const am = (height - z.T - z.B - 2 * z.P) / (ac - ab);
            const an = (width - z.L - z.R - 2 * z.P) / (af - ae);
            for (let i = 0; i < curves.length; i++) {
                const ap = curves[i];
                this.ctx.strokeStyle = ap.color;
                this.ctx.setLineDash([]);
                this.ctx.fillStyle = this.ctx.strokeStyle + "08";
                this.ctx.beginPath();
                this.ctx.lineTo(this.getX(ae, ae, an, z), this.getY(ab, ab, am, z, height));
                for (let aq = 0; aq < ap.data.length; aq++) {
                    if (ap.data[aq] !== null) {
                        this.ctx.lineTo(this.getX(ap.step[aq], ae, an, z), this.getY(ap.data[aq], ab, am, z, height));
                    }
                }
                this.ctx.lineTo(this.getX(af, ae, an, z), this.getY(ap.data[ap.data.length - 1], ab, am, z, height));
                this.ctx.stroke();
            }
            this.ctx.setLineDash([]);
            this.ctx.font = "12px Lato";
            this.ctx.fillStyle = "#EEE";
            this.ctx.textBaseline = "top";
            this.ctx.textAlign = "center";
            for (let ar = Math.floor(ae / 60) + 1; ar < Math.floor(af / 60) + 1; ar++) {
                const as = ar * 60;
                this.ctx.fillText(ar + " min", this.getX(as, ae, an, z), height - z.B + 8);
                this.ctx.fillText("|", this.getX(as, ae, an, z), height - z.B - 5);
            }
            this.ctx.textBaseline = "middle";
            this.ctx.textAlign = "right";
            const at = [5, 20, 100, 500, 2000, 10000, 50000, 200000, 1000000];
            let au = 1;
            for (let i = 0; i < at.length; i++) {
                if (ac - ab > at[i] * 4) {
                    au = at[i];
                } else {
                    break;
                }
            }
            for (let aw = Math.floor(ab / au) + 1; aw < Math.floor(ac / au) + 1; aw++) {
                const ax = aw * au;
                this.ctx.fillText(this.formatNumber(ax) + " --", z.L + 4, this.getY(ax, ab, am, z, height));
            }
            if (this.mouse.in) {
                let ay = 0;
                let az = 0;
                let ba = 0;
                let bb = -1;
                for (let i = 0; i < curves.length; i++) {
                    const curve = curves[i];
                    for (let j = 0; j < curve.step.length; j++) {
                        const bf = Math.abs(this.getX(curve.step[j], ae, an, z) - this.mouse.x);
                        if (bb === -1 || bf < bb) {
                            bb = bf;
                            ay = curve.step[j];
                            az = j;
                        }
                    }
                }
                bb = 999;
                this.ctx.strokeStyle = "#EEE";
                for (let i = 0; i < curves.length; i++) {
                    const curve = curves[i];
                    if (curve.step.length > 0) {
                        const bi = this.getLastIndex(curve, ay);
                        if (bi > -1) {
                            const bj = Math.abs(this.getY(curve.data[bi], ab, am, z, height) - this.mouse.y);
                            if (bj < bb) {
                                bb = bj;
                                ba = i;
                            }
                            if (curve.data[bi] !== null) {
                                this.ctx.beginPath();
                                this.ctx.arc(this.getX(curve.step[bi], ae, an, z), this.getY(curve.data[bi], ab, am, z, height), this.ctx.lineWidth * 1.5, 0, 2 * Math.PI);
                                this.ctx.stroke();
                            }
                        }
                    }
                }
                const bk = this.getLastIndex(curves[ba], ay);
                if (bk > -1) {
                    const bl = this.getX(curves[ba].step[bk], ae, an, z);
                    const bm = this.getY(curves[ba].data[bk], ab, am, z, height);
                    this.ctx.beginPath();
                    this.ctx.arc(bl, bm, this.ctx.lineWidth * 2.5, 0, 2 * Math.PI);
                    this.ctx.stroke();
                    this.ctx.lineWidth = 1;
                    this.ctx.beginPath();
                    this.ctx.moveTo(z.L, bm);
                    this.ctx.lineTo(width - z.R, bm);
                    this.ctx.moveTo(bl, z.T);
                    this.ctx.lineTo(bl, height - z.B);
                    this.ctx.stroke();
                }
                const bn = this.mouse.x < width / 2 ? width - 280 : z.L;
                this.ctx.font = "600 28px Arial";
                this.ctx.textBaseline = "middle";
                this.ctx.textAlign = "center";
                this.ctx.fillStyle = "rgba(64, 60, 60, 0.25)";
                this.ctx.fillRect(bn, 0, 280, height - z.B);
                this.ctx.fillStyle = "rgba(255, 255, 255, 0.8)";
                const bo = this.formatTime(curves[ba].step[bk]);
                this.ctx.fillText(bo, 150 + bn, 24);
                const bp = [];
                let bq = "";
                for (let i = 0; i < curves.length; i++) {
                    const curve = this.getLastIndex(curves[i], ay);
                    if (curve > -1) {
                        if (bt === ba) {
                            bq = curves[i].label;
                        }
                        bp.push({
                            label: curves[i].name,
                            color: curves[i].color,
                            value: curves[i].data[curve]
                        });
                    }
                }
                bp.sort(function (a, b) {
                    return a.value < b.value ? 1 : b.value < a.value ? -1 : 0;
                });
                this.ctx.font = "600 18px Arial";
                for (let i = 0; i < bp.length; i++) {
                    this.ctx.textAlign = "left";
                    this.ctx.fillStyle = bp[i].color;
                    this.ctx.fillStyle = bp[i].label === bq ? this.ctx.fillStyle + "DD" : this.ctx.fillStyle + "AA";
                    this.ctx.fillText(bp[i].label, bn + 10, 50 + i * 20);
                    this.ctx.textAlign = "right";
                    if (bp[i].value !== null) {
                        this.ctx.fillText(this.formatNumber(bp[i].value), bn + 275, 50 + i * 20);
                    } else {
                        this.ctx.fillText("error", bn + 275, 50 + i * 20);
                    }
                }
            }
        }
    },
    watch: {
        curves: function () {
            this.drawChart();
        }
    },
    computed: {},
    provide() {
        return {
            provide: this.canvas
        };
    },
    created() { },
    mounted() {
        const canvas = this.$refs["chart-canvas-" + this.index];
        this.canvas.context = canvas.getContext("2d");
        this.ctx = this.canvas.context;
        canvas.width = canvas.clientWidth;
        canvas.height = canvas.clientHeight;
        this.width = canvas.width;
        this.height = canvas.height;
        canvas.addEventListener("mouseover", ca => {
            this.mouseIn(ca);
        });
        canvas.addEventListener("mousemove", cb => {
            this.mouseMove(cb);
        });
        canvas.addEventListener("mouseout", cc => {
            this.mouseOut();
        });
        window.addEventListener("resize", cd => {
            canvas.width = canvas.clientWidth;
            canvas.height = canvas.clientHeight;
            this.width = canvas.width;
            this.height = canvas.height;
            this.drawChart();
        });
        this.drawChart();
    }
});