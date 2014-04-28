from flask import Flask, abort
import sys
app = Flask(__name__)


@app.route('/')
def root():
    return 'OK'


@app.route('/tetra/<int:val>')
def tetra(val):
    val = int(val)
    return str(val ** val)


# health check for load balancer
@app.route('/available')
def available():
    return 'yes'

if __name__ == '__main__':
    try:
        port = int(sys.argv[1])
    except IndexError:
        port = 5000
    print 'running on port %d' % port
    app.run(port=port, threaded=True)
