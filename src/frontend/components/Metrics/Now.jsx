import _ from 'lodash';

// React
import React from 'react';
import { connect } from 'react-redux';

// Actions
import { Metrics as MetricAction } from 'frontend/actions';

// Components
import { Empty, realtimeColorAccessor, realtimeColorAccessorByTime } from './Common';
import ReactEcharts from 'react-echarts-component';

class Now extends React.Component {

  componentWillMount() {
    const { dispatch } = this.props;
    dispatch(MetricAction.enterPage());
  }

  componentDidMount() {
    const { dispatch, project, realtime } = this.props;
    this.timer = setInterval(() => {
      dispatch(MetricAction.getOverviewRealtime(project.apikey, realtime));
    }, 5000);
    dispatch(MetricAction.getOverviewRealtime(project.apikey, realtime));
  }

  componentWillUnmount() {
    clearInterval(this.timer);
  }

  _renderResponseStatusPie() {
    const { realtime } = this.props;
    const dataLen = realtime[0].length;
    if (dataLen === 0) {
      return <Empty title="Current Transaction" />;
    }
    const option = {
      title: {},
      tooltip: {
        trigger: 'item',
        formatter: 'Response time {b}<br/>{c} ({d}%)',
      },
      legend: {
        orient: 'vertical',
        x: 'right',
        data: [],
      },
      series: [],
    };
    const tempData = {
      type: 'pie',
      radius: ['50%', '70%'],
      avoidLabelOverlap: false,
      label: {
        normal: {
          show: false,
        },
        emphasis: {
          show: false,
        },
      },
      labelLine: {
        normal: {
          show: false,
        },
      },
      data: [],
    };
    let count = 0;
    const grouped = {};
    for (let i = 0; i < dataLen; i++) {
      const tg = parseInt(realtime[0][i].timegroup, 10);
      if (realtime[0][i].sum === null) {
        continue;
      }
      count += realtime[0][i].sum;
      if (tg <= 12) {
        if (!grouped['0']) {
          grouped['0'] = 0;
        }
        grouped['0'] += realtime[0][i].sum;
      } else if (tg <= 32) {
        if (!grouped['1']) {
          grouped['1'] = 0;
        }
        grouped['1'] += realtime[0][i].sum;
      } else {
        if (!grouped['2']) {
          grouped['2'] = 0;
        }
        grouped['2'] += realtime[0][i].sum;
      }
    }
    option.title.text = `${count} Transaction(s)`;
    _.forEach(grouped, (value, k) => {
      let name = '';
      if (k === '0') {
        name = 'under 3s';
        option.legend.data.push('under 3s');
      } else if (k === '1') {
        name = 'under 8s';
        option.legend.data.push('under 8s');
      } else {
        name = 'over 8s';
        option.legend.data.push('over 8s');
      }
      tempData.data.push({
        name,
        value,
        itemStyle: {
          normal: {
            color: realtimeColorAccessor(k),
          },
        },
      });
    });
    option.series.push(tempData);
    option.legend.data = _.sortBy(option.legend.data);
    return (
      <div>
        <div className="chart-wrapper-header">Transaction (last 5 second)</div>
        <div className="chart-wrapper">
          <ReactEcharts option={option} height={300} />
        </div>
      </div>
    );
  }

  _renderResponseScatter() {
    const { realtime } = this.props;
    const dataLen = realtime[1].length;
    if (dataLen === 0) {
      return <Empty title="Hitmap" />;
    }
    const option = {
      tooltip: {
        formatter: (params) => {
          return `${params.value[3]}<br/>${params.value[0]}<br/>${params.value[2]} transaction(s)`;
        },
      },
      legend: {
        data: [],
        left: 'right',
      },
      xAxis: [
        {
          type: 'time',
          scale: true,
          splitLine: {
            show: false,
          },
        },
      ],
      yAxis: [
        {
          type: 'value',
          scale: true,
          axisLabel: {
            formatter: '{value}ms',
          },
          splitLine: {
            lineStyle: {
              type: 'dotted',
            },
          },
        },
      ],
      series: [],
    };
    const grouped = {};
    for (let i = 0; i < dataLen; i++) {
      const tg = parseInt(realtime[1][i].timegroup, 10);
      if (realtime[1][i].sum === null) {
        continue;
      }
      if (!grouped[tg]) {
        grouped[tg] = [];
      }
      grouped[tg].push(realtime[1][i]);
    }
    const tempData = {};
    _.forEach(grouped, (value, k) => {
      const name = `${k * 250}ms`;
      tempData[name] = {
        name,
        type: 'scatter',
        itemStyle: {
          normal: {
            color: realtimeColorAccessorByTime(k),
          },
        },
        data: [],
      };
      _.forEach(value, (d) => {
        tempData[name].data.push([
          new Date(d.time),
          k * 250,
          d.sum,
          `${((k === 0 ? 1 : k) - 1) * 250} ~ ${k * 250}ms`]);
      });
    });
    _.forEach(tempData, (v) => {
      option.series.push(v);
    });
    return (
      <div>
        <div className="chart-wrapper-header">Hitmap</div>
        <div className="chart-wrapper">
          <ReactEcharts option={option} height={300} />
        </div>
      </div>
    );
  }

  render() {
    return (
      <div>
        {this._renderResponseStatusPie()}
        {this._renderResponseScatter()}
      </div>
    );
  }
}

Now.propTypes = {
  dispatch: React.PropTypes.func.isRequired,
  project: React.PropTypes.object,
  realtime: React.PropTypes.array,
};

const mapStateToProps = (state) => ({
  project: state.project.project.data,
  realtime: state.metrics.overview.realtime.data,
});

export default connect(mapStateToProps)(Now);