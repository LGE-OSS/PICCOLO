/**
*  Copyright(c) 2022 ZettaScale Technology and others
*
*   This program and the accompanying materials are made available under the
*   terms of the Eclipse Public License v. 2.0 which is available at
*   http://www.eclipse.org/legal/epl-2.0, or the Eclipse Distribution License
*   v. 1.0 which is available at
*   http://www.eclipse.org/org/documents/edl-v10.php.
*
*   SPDX-License-Identifier: EPL-2.0 OR BSD-3-Clause
*/

#include <iomanip>
#include <iostream>
#include <fstream>
#include <thread>
#include <chrono>

#include "dds/dds.h"
#include "dds/dds.hpp"
#include "RoundTrip.hpp"

using namespace std;
using namespace org::eclipse::cyclonedds;

static bool D = true;

int main (int argc, char *argv[])
{
    dds::domain::DomainParticipant participant(domain::default_id());
    dds::topic::Topic<RoundTripModule::DataType> topic1(participant, "ping");
    dds::topic::Topic<RoundTripModule::DataType> topic2(participant, "pong");

    dds::sub::Subscriber subscriber(participant);
    dds::sub::DataReader<RoundTripModule::DataType> reader(subscriber, topic1);

    dds::pub::Publisher publisher(participant);
    dds::pub::DataWriter<RoundTripModule::DataType> writer(publisher, topic2);

    cout << "=== [Subscriber] Waiting for publisher. [doamin_id] " << participant->domain_id() << endl;

    while (writer.publication_matched_status().current_count() == 0) {
        this_thread::sleep_for(chrono::milliseconds(20));
    }

    cout << "=== [Subscriber] success finding publisher. " << endl;

    this_thread::sleep_for(chrono::milliseconds(1000));
    while (true) {
        auto samples = reader.take();
        if (samples.length() > 0) {
            writer.write(samples.begin()->data());
            this_thread::sleep_for(chrono::milliseconds(500));
        }
    }

    return EXIT_SUCCESS;
}

