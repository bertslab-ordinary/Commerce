# Use an official Jupyter image as the base image
FROM jupyter/datascience-notebook:latest

# Set the working directory
WORKDIR /home/jovyan

# Install Go
RUN conda install -c anaconda go -y

# Set the GOPATH and GOBIN environment variables
ENV GOPATH /home/jovyan/go
ENV GOBIN /home/jovyan/go/bin
ENV PATH $GOPATH/bin:$PATH

# Install necessary Python packages
RUN pip install --upgrade pip
RUN pip install jupyterlab-go-to-definition

# Expose the JupyterLab port
EXPOSE 8888

# Start JupyterLab
CMD ["jupyter", "lab", "--no-browser", "--ip=0.0.0.0", "--port=8888", "--allow-root", "--NotebookApp.token=''"]