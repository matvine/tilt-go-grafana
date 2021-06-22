# -*- mode: Python -*-
load(
    'ext://coreos_prometheus',
    'setup_monitoring',
    'get_prometheus_resources',
    'get_prometheus_dependencies',
)

setup_monitoring()

load (
    'ext://configmap', 
    'configmap_create'
)

configmap_create('tilt-grafana-config',
                 from_file=[
                   'grafana.ini=./monitoring/grafana/grafana.ini',
                   'datasource-prometheus.yaml=./monitoring/grafana/datasource-prometheus.yaml',
                   ]
)

docker_build('tilt-go-grafana-image', '.', dockerfile='deployments/Dockerfile')

k8s_yaml('deployments/application.yaml')
k8s_yaml('monitoring/grafana/grafana.yaml')
k8s_yaml('monitoring/loki/loki.yaml')
k8s_yaml('monitoring/promtail/promtail.yaml')


k8s_resource('app',
    objects=get_prometheus_resources('deployments/application.yaml', 'app'),
    resource_deps=get_prometheus_dependencies(),
    port_forwards=8080
)

k8s_resource('grafana-server',
             port_forwards=[
               port_forward(10354, 3000, name='grafana'),
             ],
             resource_deps=['app']
)

k8s_resource('loki')
k8s_resource('promtail')