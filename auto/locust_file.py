import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def request_home(self):
        self.client.get("/")

    @task
    def request_index(self):
        self.client.get("/index")