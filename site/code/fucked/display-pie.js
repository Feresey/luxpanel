Vue.component("display-pie", {
    props: {
        index: {
            type: Number,
            required: true
        },
        data: {
            type: Array,
            required: true
        }
    },
    template: "<div class=\"full\">\n\n\t\t<canvas class=\"full\" :ref=\"'pie-canvas-'+index\">Canvas not enabled</canvas>\n\t\n\t</div>",
    data() {
        return {
            canvas: {
                context: null
            },
            ctx: null,
            size: 0,
            p: 0,
            cx: 0,
            cy: 0
        };
    },
    methods: {
        getPoint(g) {
            if (g < 0.25) {
                return {
                    "x": -1 + g * 8,
                    "y": -1
                };
            } else if (g < 0.5) {
                return {
                    "x": 1,
                    "y": -3 + g * 8
                };
            } else if (g < 0.75) {
                return {
                    "x": 5 - g * 8,
                    "y": 1
                };
            } else {
                return {
                    "x": -1,
                    "y": 7 - g * 8
                };
            }
        },
        draw() {
            const h = this.p;
            this.ctx.fillStyle = '#000';
            this.ctx.fillRect(0, 0, this.size, this.size);
            if (!this.data || !this.data.length) {
                return;
            }
            this.ctx.textAlign = 'center';
            this.ctx.textBaseline = 'middle';
            this.ctx.font = "" + Math.round(2.5 * h) + "px Arial";
            let i = 0;
            for (let j = 1; j < this.data.length; j++) {
                const k = this.data[j];
                const l = k.amount / this.data[0].amount;
                const m = this.getPoint(i);
                const o = this.getPoint(i + l);
                this.ctx.beginPath();
                this.ctx.moveTo(this.cx, this.cy);
                this.ctx.lineTo(this.cx + m.x * 50 * h, this.cy + m.y * 50 * h);
                for (let q = 0; q < 4; q++) {
                    if (Math.floor(i * 4) + q < Math.floor((i + l) * 4)) {
                        const r = this.getPoint(Math.floor(i * 4 + 1 + q) / 4);
                        this.ctx.lineTo(this.cx + r.x * 50 * h, this.cy + r.y * 50 * h);
                    } else {
                        break;
                    }
                }
                this.ctx.lineTo(this.cx + o.x * 50 * h, this.cy + o.y * 50 * h);
                this.ctx.closePath();
                this.ctx.fillStyle = this.getColor(j);
                this.ctx.fillStyle = this.ctx.fillStyle + "DD";
                this.ctx.fill();
                i += l;
            }
            i = 0;
            for (let s = 1; s < this.data.length; s++) {
                const elem = this.data[s];
                const value = elem.amount / this.data[0].amount;
                const w = this.getPoint(i + value / 2);
                const z = {
                    "x": this.cx + w.x * 35 * h,
                    "y": this.cy + w.y * 35 * h
                };
                this.ctx.fillStyle = "rgba(0,0,0,0.5)";
                this.ctx.fillRect(z.x - 8 * h, z.y - 5 * h, 16 * h, 10 * h);
                this.ctx.fillStyle = '#FFF';
                this.ctx.fillText(elem.name, z.x, z.y - 3 * h);
                this.ctx.fillText(this.formatNumber(elem.amount), z.x, z.y);
                this.ctx.fillText(Math.round(value * 100) + " %", z.x, z.y + 3 * h);
                i += value;
            }
            this.ctx.fillStyle = '#000';
            this.ctx.fillRect(this.cx - 10 * h, this.cy - 10 * h, 20 * h, 20 * h);
            this.ctx.fillStyle = '#FFF';
            this.ctx.fillText("Total " + this.data[0].name, this.cx, this.cy - 6 * h);
            this.ctx.fillText(this.formatNumber(this.data[0].amount), this.cx, this.cy - 2 * h);
            this.ctx.fillText("Average", this.cx, this.cy + 2 * h);
            this.ctx.fillText(this.formatNumber(this.data[0].average), this.cx, this.cy + 6 * h);
        }
    },
    watch: {
        data: function () {
            this.draw();
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
        const elem = this["$refs"]['pie-canvas-' + this.index];
        this.canvas.context = elem.getContext("2d");
        this.ctx = this.canvas.context;
        elem.width = elem.parentElement.clientWidth;
        elem.height = elem.parentElement.clientHeight;
        this.size = Math.min(elem.width, elem.height);
        this.p = this.size / 100;
        this.cx = this["$refs"]['pie-canvas-' + this.index].width / 2;
        this.cy = this["$refs"]["pie-canvas-" + this.index].height / 2;
        this.draw();
    }
});