from django.conf.urls import include, url
from pokemon_v3 import urls as pokemon_v3_urls

# pylint: disable=invalid-name

urlpatterns = [
    url(r"^", include(pokemon_v3_urls)),
]
