FROM  atricore/josso-ee:2.6.0-unstable

COPY ./etc/atricore.lic /opt/atricore/josso-ee/server/etc/atricore.lic

COPY ./etc/host.key /opt/atricore/josso-ee/server/etc/host.key

COPY ./etc/org.atricore.idbus.kernel.main.cfg /opt/atricore/josso-ee/server/etc/org.atricore.idbus.kernel.main.cfg

RUN chown -R atricore /opt/atricore

RUN chgrp -R atricore /opt/atricore