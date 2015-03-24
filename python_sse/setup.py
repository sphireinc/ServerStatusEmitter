from setuptools import setup
import sys

setup(
    name="serverstatusemitter",
    version="1.0-dev",
    install_requires=[
        "psutils",
        "requests>=0.8.8",
    ],
    tests_require=["tox"],
    packages=[
        'serverstatusemitter',
    ],
    author="Juan L. Sanchez",
    author_email="juan.sanchez@juanleonardosanchez.com",
    url="https://bitbucket.org/sphire-development/serverstatusemitter",
    description="The Server Status Monitoring Reporter",
    entry_points={ },
)
