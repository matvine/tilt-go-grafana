# Load Test
`locust --headless -f locust_file.py --only-summary -u 10 -r 20 -t 2m --stop-timeout 10 -H http://localhost:8080`