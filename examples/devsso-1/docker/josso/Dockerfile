FROM  atricore/josso-ee:2.6.0-unstable

COPY ./atricore.lic /opt/atricore/josso-ee/server/etc/atricore.lic

COPY host.key /opt/atricore/josso-ee/server/etc/host.key

RUN chown -R atricore /opt/atricore

RUN chgrp -R atricore /opt/atricore