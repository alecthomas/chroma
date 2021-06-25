#########################################################################
# This is an example VCL file for Varnish 4.0.				#
# From: https://gist.github.com/davidthingsaker/6b0997b641fdd370a395    #
# LICENSE: If this could help you in any way, you are obliged to use it	#
# for free with no limitations. 					#
#########################################################################


# Marker to tell the VCL compiler that this VCL has been adapted to the
# new 4.0 format.
vcl 4.0;

import std;

# Default backend definition. Set this to point to your content server.
backend default {
    .host = "127.0.0.1";
    .port = "8080";
}

