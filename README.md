<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/kyledinh/protoserver-go">
    <img src="https://gitlab.com/uploads/-/system/project/avatar/16184934/icons8-Octopus-96.png?width=64" alt="Logo" width="64" height="64">
  </a>

  <h3 align="center">Protoserver</h3>

  <p align="center">
    Configurable containers to prototype services for a Kubernetes stack!
    <br />
    <!-- TODO: do documentation
    <a href="https://github.com/kyledinh/protoserver-go"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/kyledinh/protoserver-go">View Demo</a>
    ·
    <a href="https://github.com/kyledinh/protoserver-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/kyledinh/protoserver-go/issues">Request Feature</a>
    -->
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li>
      <a href="#this-repo-dvelopement">This Repo Development</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>
<br/>

<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

Rapid prototyping can be a great tool for desinging and developing a project. Getting a Kubernetes stack up to test the flow of an application is great when each service will included structured logging and monitoring services.

Each Protoserver can be configured to certain MODES:

* STUB - to return a fixed response for each endpoint 
* RELAY - to forward request to another internal or external service  
* AGENT - to process request based on configured preset MACROS 

You can progressively replace mocked/stubbed services with your own services. This project will let you start designing at the Kubernetes stack level.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

This section should list any major frameworks/libraries used to bootstrap your project. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.

* [![Go][Golang]][Go-url]
* [![Docker]][Docker-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

Usage of this respo will soon only require pre-built Docker Images that will be on  https://hub.docker.com/u/kyledinh. I will be writing Usage Documentation soon, the work so far can be found in the `examples` directory.

<!--
_For more examples, please refer to the [Documentation](https://example.com)_
-->

See the `examples` directory to get started.
- [ ] Provide Docker container images on: https://hub.docker.com/u/kyledinh 
- [ ] Examples to deploy locally with Docker Desktop 
- [ ] Examples to deploy hosted Kubernetes stack like Linode or GKE
- [ ] Provide tools to view/manage Kubernetes stack (kt)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- DEVELOPMENT GETTING STARTED -->
## This Repo Development 

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

| Software       | Version | Install                                        |
|----------------|---------|------------------------------------------------|
| Go             | 1.18    | https://go.dev/doc/install                     |
| Docker Desktop | 4.3.x   | https://www.docker.com/products/docker-desktop |

This script will query your local machine for software and their version.
* Check your local installation for dependencies and their version 
  ```sh
  make check
  ```

### Installation

_The Makefile will provide scripts to install all the local dependencies._

1. Clone the repo
   ```sh
   git clone https://github.com/kyledinh/protoserver-go.git
   ```
2. Install packages
   ```sh
   make setup 
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>




<!-- ROADMAP -->
## Roadmap

- [x] STUB, RELAY action 
- [x] Structured Logging 
- [ ] Auth Middleware with JWT 
- [ ] Tooling for development - linkt to https://github.com/kyledinh/btk-go
- [ ] Tooling Kubernetes stack, deploy/view/monitoring    
- [ ] Prometheus monitoring 
- [ ] Example Usages 
    - [ ] Local 
    - [ ] GKE

See the [open issues](https://github.com/kyledinh/protoserver-go/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the Apache License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Kyle Dinh - [@iamslowblood](https://twitter.com/iamslowblood) 

Project Link: [https://github.com/kyledinh/protoserver-go](https://github.com/kyledinh/protoserver-go)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Best README Template](https://github.com/othneildrew/Best-README-Template)
* [BTK - Go Project Dev Tools](https://github.com/kyledinh/btk-go)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png
[Go-url]: https://go.dev
[Docker-url]: https://www.docker.com/