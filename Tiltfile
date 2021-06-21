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
                   'grafana.ini=./monitoring/grafana.ini',
                   'datasource-prometheus.yaml=./monitoring/datasource-prometheus.yaml',
                   ]
)

docker_build('tilt-go-grafana-image', '.', dockerfile='deployments/Dockerfile')

k8s_yaml('deployments/kubernetes.yaml')
k8s_yaml('deployments/prometheus.yaml')
k8s_yaml('monitoring/grafana.yaml')

k8s_resource('tilt-go-grafana-app',
    objects=get_prometheus_resources('deployments/kubernetes.yaml', 'tilt-go-grafana'),
    resource_deps=get_prometheus_dependencies(),
    port_forwards=8080
)

k8s_resource('tilt-go-grafana-server',
             port_forwards=[
               port_forward(10354, 3000, name='grafana'),
             ],
             resource_deps=['tilt-go-grafana-app']
)